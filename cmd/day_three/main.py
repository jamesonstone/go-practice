import asyncio
import os
import json
import aiofiles

class MemDb:
  def __init__(self):
    self._lock = asyncio.Lock()
    self._data = {}

  async def get(self, key):
    """
    key value from key
    """
    async with self._lock:
      return self._data.get(key)


  async def set(self, key, value):
    async with self._lock:
      if key not in self._data:
        self._data[key] = value
        return True
      return False

  async def get_all(self):
    async with self._lock:
      return dict(self._data) # return a copy of the dict

  async def clear_all(self):
    async with self._lock:
      self._data = {}
      return self._data

  async def filter(self, predicate):
    async with self._lock:
      # {key_expression: value_expression for item in iterable if condition}
      return {k: v for k, v in self._data.items() if predicate(k, v)} # <--- need to note the structure

  async def save_to_disk(self, filename="memdb_data.json"):
      async with self._lock:
          try:
              # 1. Create directories if needed
              os.makedirs(os.path.dirname(filename) or ".", exist_ok=True)

              # 2. Convert data to a serializable format
              serializable_data = {}
              for k, v in self._data.items():
                key = str(k) if not isinstance(k, str) else k
                serializable_data[key] = v

              # 3. Write to file safely
              tmp_filename = filename + '.tmp'
              with open(tmp_filename, 'w') as f:
                json.dump(serializable_data, f)
              os.replace(tmp_filename, filename) # <-- atomic action

              # 4. Return success indication
              return True
          except Exception as e:
            return False
              # Handle errors gracefully


  async def restore_from_disk(self, filename):
    async with self._lock:
      try:
        # 1. check for filename on local disk
        if not os.path.exists(filename):
          return False

        # 2. json.loads the contents and deserializing
        with open(filename, 'r') as f:
          restored_data = json.load(f)

        for k, v in restored_data.items():
          if k.isdigit():
            k = int(k)
          self._data[k] = v

        return True
      except Exception as e:
        print(f"exception found: {e}")
        return False

  async def save_to_disk_streaming(self, filename):
    async with self._lock:
      try:
        os.makedirs(os.path.dirname(filename) or '.', exist_ok=True)
        async with aiofiles.open(filename, 'w') as f:
          for k, v in self._data.items():
            key = str(k) if not isinstance(k, str) else k
            record = {"key": key, "value": v}
            line = json.dumps(record)
            await f.write(line + '\n')
          return True
      except Exception as e:
        return False

  async def restore_from_disk_streaming(self, filename):
    async with self._lock:
      try:
        if not os.path.exists(filename):
          return False

        async with aiofiles.open(filename, 'r') as f:
          async for line in f:
            record = json.loads(line.strip())
            k, v = record["key"], record["value"]
            self._data[k]= v

        return True


      except Exception as e:
        print(f"exception thrown: {e}")
        return False




async def main():
  mem_db = MemDb()

  start = await mem_db.get_all()
  print(f"start: {start}")

  a = await mem_db.get("test")
  print(f"a: {a}")

  await mem_db.set("a1", 1)
  await mem_db.set("a2", 2)
  await mem_db.set(1, "a3")

  b = await mem_db.get_all()
  print(f"b: {b}")

  e = await mem_db.filter(lambda k, v: isinstance(v, int) and v > 1) # <--- need to note the structure
  print(f"e: {e}")

  # f = await mem_db.save_to_disk("snapshot1.json")
  # print(f"f: {f}")

  # g = await mem_db.clear_all()
  # print(f"g {g}")

  # h = await mem_db.restore_from_disk("snapshot1.json")
  # print(f"h {h}")

  i = await mem_db.get_all()
  print(f"i: {i}")

  j = await mem_db.clear_all()
  print(f"j {j}")

  await mem_db.set("a2", 2)
  await mem_db.set(1, "a3")

  k = await mem_db.save_to_disk_streaming("streaming.json")
  print(f"k {k}")

  await mem_db.set("a2", 2)
  await mem_db.set(1, "a3")

  m = await mem_db.get_all()
  print(f"m: {m}")

  l = await mem_db.restore_from_disk_streaming("streaming.json")
  print(f"l, {l}")

  j = await mem_db.get_all()
  print(f"j: {j}")

if __name__ == '__main__':
  asyncio.run(main())

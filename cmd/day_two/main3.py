import asyncio
import json
import os

class MemDb():
  def __init__(self):
    self._lock = asyncio.Lock()
    self._data = {}

  async def get(self, key):
    return self. _data.get(key) # returns None if key doesn't exist

  async def set(self, key, value):
    async with self._lock: # lock here to ensure we're not overwriting
      self._data[key] = value
      return True

  async def delete(self, key):
    async with self._lock:
      if key in self._data: # look before leap; faster when keys are missing often; try/except is slower exception handling overhead
        del self._data[key]
        return True
      return False

  async def filter(self, predicate):
    """
    filter items based on predicate function
    args:
      predicate: a function that takes key/value, returns True to keep
      ex. lambda k, v: isinstance(v, int) and v > 10
    returns dictionary with filtered items
    """
    async with self._lock:
      return {k: v for k, v in self._data.items() if predicate(k, v)}

  async def get_all(self):
    async with self._lock:
      return {k:v for k, v in self._data.items()}


  async def clear(self):
    async with self._lock:
      self._data = {}
      return True


  async def save_to_disk(self, filename = "memdb_data_.json"):
    async with self._lock:
      try:
        os.makedirs(os.path.dirname(filename) or ".", exist_ok = True) # create the file

        serializable_data = {}

        for k, v in self._data.items():
          key = str(k) if not isinstance(k, str) else k
          serializable_data[key] = v

        temp_filename = filename + '.tmp'
        with open(temp_filename, 'w') as f:
          json.dump(serializable_data, f)

        os.replace(temp_filename, filename)
        return True
      except Exception as e:
        print(f"Error saving to disk {e}")
        return False

  async def restore_from_disk(self, filename):
    if not os.path.exists(filename):
      return False

    try:
      with open(filename, 'r') as f:
        loaded_data = json.load(f)

      restored_data = {}
      for k, v in loaded_data.items():
        try:
          if k.isdigit():
            restore_key = int(k)
          else:
            restore_key = k
        except:
          restore_key = k
        restored_data[restore_key] = v
      async with self._lock:
        self._data = restored_data
      return True
    except Exception as e:
      print(f"error restoring from disk: {e}")
      return False




async def main():
  print("running!")
  mdb = MemDb()
  a = await mdb.get("test")
  print(f"a: {a}")
  await mdb.set("test", "this is a test value")
  b = await mdb.get("test")
  print(f"b: {b}")
  await mdb.set("t1", 1)
  c = await mdb.get("t1")
  print(f"c: {c}")
  await mdb.delete("t1")
  d = await mdb.get("t1")
  print(f"d: {d}")
  await mdb.delete("t1")
  await mdb.set(1, 1)
  await mdb.set(2, 2)
  await mdb.set(3, 3)
  await mdb.set(4, 4)
  await mdb.set("potato", "snack")
  e = await mdb.filter(lambda k, v: v == 1)
  print(f"e {e}")
  f = await mdb.get_all()
  print(f"f: {f}")
  await mdb.save_to_disk('test-save.json')
  await mdb.clear()
  g = await mdb.get_all()
  print(f"g: {g}")
  await mdb.restore_from_disk('test-save.json')
  h = await mdb.get_all()
  print(f"h: {h}")






if __name__ == '__main__':
  asyncio.run(main())

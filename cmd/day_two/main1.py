import asyncio


lock = asyncio.Lock()

async def worker(name, sleepTimer):
  # async with lock:
  print(f"{name} 🔒")
  await asyncio.sleep(sleepTimer)
  print(f"{name} 👐")



async def main():
  await asyncio.gather(
    worker("worker1", 1),
    worker("worker2", 2),
    worker("worker3", 3),
    worker("worker4", 2),
    worker("worker5", 1),
    worker("worker6", 3),
    worker("worker7", 2),
    worker("worker8", 1),
    worker("worker9", 5),
    worker("worker10", 7)
  )


if __name__ == '__main__':
 asyncio.run(main())

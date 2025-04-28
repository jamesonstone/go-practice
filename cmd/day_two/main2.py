import asyncio
import aiohttp

semaphore = asyncio.Semaphore(5)

urls = [
    "https://example.com",
    "https://httpbin.org/delay/1",
    "https://httpbin.org/delay/2",
    "https://httpbin.org/delay/3",
    "https://httpbin.org/status/200",
    "https://httpbin.org/status/404",
    "https://httpbin.org/status/500",
    "https://example.com",
    "https://httpbin.org/delay/1",
    "https://httpbin.org/delay/2"
]


async def fetch(session, url, name):
  async with semaphore:
    print(f"{name} fetching with url: {url} 🏃")
    try:
      async with session.get(url, timeout=5) as response:
        await response.text()
        print(f"{name} found: {response.status} 👊")
    except asyncio.TimeoutError:
      print(f"{name} for {url} timed out 🕐")
    except Exception as e:
      print(f"{name} for url: {url} with error: {str(e)} ❗")


async def main():
  async with aiohttp.ClientSession() as session:
    async with asyncio.TaskGroup() as tg:
      for i, url in enumerate(urls):
        tg.create_task(fetch(session, url, f"worker{i+1}"))
  pass



if __name__ == '__main__':
  asyncio.run(main())

import requests
import threading
import sys

url = "http://faasgray1.thinktan.site/info"
concurrency = int(sys.argv[1]) if len(sys.argv) > 1 else 10

results = []

def worker(i):
    try:
        resp = requests.get(
            url,
            timeout=5,
            verify=False,
            proxies={"http": None, "https": None}  # 显式禁用代理
        )
        print(f"[{i}] Status: {resp.status_code}")
        results.append(resp.status_code)
    except Exception as e:
        print(f"[{i}] Error: {e}")
        results.append("ERR")

threads = []
for i in range(concurrency):
    t = threading.Thread(target=worker, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

# 汇总统计
print("\n=== Summary ===")
from collections import Counter
summary = Counter(results)
for code, count in summary.items():
    print(f"{code}: {count}")

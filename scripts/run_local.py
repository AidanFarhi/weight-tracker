import os
import signal
import subprocess

env = os.environ.copy()

env["WT_APP_ENV"] = "local"

proc = subprocess.Popen(["air"], env=env)

try:
    proc.wait()
except KeyboardInterrupt:
    proc.send_signal(signal.SIGINT)
    proc.wait()

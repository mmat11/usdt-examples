from time import sleep
import uuid

import stapsdt

provider = stapsdt.Provider("pyapp")
probe = provider.add_probe("pyprobe", stapsdt.ArgTypes.uint64, stapsdt.ArgTypes.int32)
provider.load()


counter = 0
while True:
    string = str(uuid.uuid4())
    probe.fire(string, counter)
    print(f"TRACEE ({counter}): fired {string}")
    counter += 1
    sleep(0.5)

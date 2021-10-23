import logging
import uuid
from time import sleep

import stapsdt

logging.basicConfig(level=logging.INFO)

provider = stapsdt.Provider("pyapp")
probe = provider.add_probe("pyprobe", stapsdt.ArgTypes.uint64, stapsdt.ArgTypes.int32)
provider.load()


if __name__ == "__main__":
    c = 0
    while True:
        uuidstr = str(uuid.uuid4())
        probe.fire(uuidstr, c)
        logging.info(f"counter={c}, uuid={uuidstr}")
        sleep(0.5)
        c += 1

# -*- coding: utf-8 -*-
# A basic converter for cs units
# useful for gauging download and upload speeds

import argparse

class CSUnitConverter():
    def __init__(self):
        self.bitSize = 1 # bits
        self.kilobitSize = 1000 # bits
        self.megabitSize = 100_000 # bits
        self.byteSizeInBits = 8 # 8 bits
        self.kilobyteSizeInBits = 8000 # 8000 bits
        self.megabyteSizeInBits = 800_000

    def getDownloadWaitTime(self, dlspeed, fsize, dlspeedUnits="Mbps", fSizeUnitBytes="MB"):
        if fSizeUnitBytes == "B":
            fsizeBits = fsize * self.byteSizeInBits
        if fSizeUnitBytes == "KB":
            fsizeBits = fsize * self.kilobyteSizeInBits
        if fSizeUnitBytes == "MB":
            fsizeBits = fsize * self.megabyteSizeInBits
        if dlspeedUnits == "Mbps":
            dlspeed = dlspeed * self.megabitSize
        return fsizeBits / dlspeed

    def getUploadWaitTime(self, uplspeed, fsize, uplspeedUnits="Mbps", fSizeUnitBytes="MB"):
        if fSizeUnitBytes == "B":
            fsizeBits = fsize * self.byteSizeInBits
        if fSizeUnitBytes == "KB":
            fsizeBits = fsize * self.kilobyteSizeInBits
        if fSizeUnitBytes == "MB":
            fsizeBits = fsize * self.megabyteSizeInBits
        if uplspeedUnits == "Mbps":
            uplspeed = uplspeed * self.megabitSize
        return fsizeBits / uplspeed

if __name__ == "__main__":


    GB = 1_000_000_000  # in bytes
    MB = 1_000_000  # in bytes
    KB = 1_000  # in bytes
    MSGSIZE = None # in bytes

    converter = CSUnitConverter()

    parser = argparse.ArgumentParser(description='A very basic units converter for cs apps. Must enter only one file size in Gibabytes, Megabytes, or Kilobytes. If multiple file sizes are entered, the last one entered will be converted only.')

    parser.add_argument('-dlspeed', metavar='Download speed in bps', type=float, default=90, help='Download speed in bits per second')

    parser.add_argument('-ulspeed', metavar='Upload speed in bps', type=float, default=40, help='Upload speed in bits per second')

    parser.add_argument('-gb', metavar='Size of file in GB', type=float, default=None, help='Gibabytes')

    parser.add_argument('-mb', metavar='Size of file in MB', type=float, default=None, help='Megabytes')

    parser.add_argument('-kb', metavar='Size of file in KB', type=float, default=None, help='Kilobytes')

    args = parser.parse_args()
    dlspeed = args.dlspeed
    ulspeed = args.ulspeed

    # test condition 
    assert any([args.gb, args.mb, args.kb]), "Must supply file size in GB, MB, or KBs"

    if args.gb:
        MSGSIZE = args.gb * GB
    if args.mb:
        MSGSIZE = args.mb * MB
    if args.kb:
        MSGSIZE = args.kb * KB

    waittime = converter.getDownloadWaitTime(dlspeed=dlspeed, fsize=MSGSIZE, fSizeUnitBytes="B")

    print("The download will take approx {} seconds.".format(waittime))

    waittime = converter.getUploadWaitTime(uplspeed=ulspeed, fsize=MSGSIZE, fSizeUnitBytes="B")
    print("The upload will take approx {} seconds.".format(waittime))

    #print("How long will it take a file of size {}{} to saturate {} of the network nodes?".format(1024,"B","80%"))

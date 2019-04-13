
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
    converter = CSUnitConverter()
    GB = 1.024e+12

    MSGSIZE =  # Bytes
    waittime = converter.getDownloadWaitTime(dlspeed=100, fsize=MSGSIZE, fSizeUnitBytes="B")

    print("The download will take approx {} seconds.".format(waittime))

    waittime = converter.getUploadWaitTime(uplspeed=10, fsize=100, fSizeUnitBytes="B")
    print("The upload will take approx {} seconds.".format(waittime))

    print("How long will it take a file of size {}{} to saturate {} of the network nodes?".format(1024,"B","80%"))

import argparse
def printHelloWorld(x,y):
  return "{} {}".format(x,y)
def testShyness(x,y):
  assert printHelloWorld(x,y)=="Hello world!", "Don't be shy, say Hello World!"

def testShyness(x,y):
  assert printHelloWorld(x,y)=="Hello world!", "Don't be shy, say Hello World!"

if __name__ == '__main__':
  parser = argparse.ArgumentParser(description='World to the homepage of Daniel Cook.')
  parser.add_argument('-x', metavar='x', type=str, default='Hello', help='Python is my first programming language.')
  parser.add_argument('-y', metavar='y', type=str, default='world!', help='This is a silly program I wrote.')
  args = parser.parse_args()
  testShyness(args.x,args.y)
  # I started programming in Python during my Master's degree. 
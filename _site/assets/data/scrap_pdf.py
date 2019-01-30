# -*- coding: utf-8 -*-
# author: DCook
# data August 2017
"""
Script to return the text on a single page of a pdf
two arguments at command line:
the full path to the pdf
the page in the pdf
writes a text file with the date (probably wanna change that if you use it a lot in a single day)
It's just a start to something larger -- getting all of the predicate-arguments from a text
currently working on Chomsky and want to get his claims in proposition form: A is B. If A is B, then Q. etc.
"""
import datetime as dt
import os
import PyPDF2 as pdf
import sys

path =  os.sep.join(os.path.normpath(os.path.abspath(__file__)).split(os.sep)[:-1])
today = dt.datetime.now()
thedate = str(today.strftime('%m-%d-%Y'))
theday =  str(today.strftime('%A'))

pdfFile = sys.argv[1]
pdfPage = int(sys.argv[2])
pdfReader = pdf.PdfFileReader(pdfFile)
text = pdfReader.getPage(pdfPage).extractText()[0:]

with open(path+os.sep+str(thedate)+"scrap.txt", "a") as textfile:
    textfile.write("<br>\n")
    textfile.write("\n <br>"+text.encode('utf8')+" <br>\n")

"""
download the pdfs
list them
open them:
pdfFile = open('APNE-04.pdf', 'rb')
read them:
extract some text:
scrap the webpages with beautiful soup
"""

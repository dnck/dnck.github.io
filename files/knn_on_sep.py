# -*- coding: utf-8 -*-
# dnck
# April 1 2018

import BeautifulSoup as bs
import requests
import os
import re
import pandas
import numpy
import graphlab
import networkx as nx
import pygraphviz as pgv
from matplotlib import pyplot as plt
from nxpd import draw, nxpdParams


#-----------@student.hu-berlin.de
#graphlab.product_key.get_product_key()

##########################################################
# Set your path. You write text files here and you save the SFrame.
workDir=os.path.normpath('/Users/yourname/ml/stanfordMine/')
# Do not be a square. Failure to follow the rules may result in data loss.
##########################################################

def scrapStanford(nruns=int):
    nruns=nruns
    for x in range(nruns):
        #start scrapping the encyclopedia
        http = requests.get("https://plato.stanford.edu/cgi-bin/encyclopedia/random")
        soup = bs.BeautifulSoup(http.content)

        #Get the p tags and string them.
        thisEntry=[]
        for tag in soup.findAll("p"):
            thisEntry.append(tag.text)
        try:
            text = [i.strip().replace("\n"," ") for i in thisEntry]
        except:
            text = [i.encode('utf-8').strip().replace("\n"," ") for i in thisEntry]

        strText=[]
        for i in text:
            try:
                i = str(i).replace("<em>"," ").replace("</em>", " ")
                strText.append(i)
            except:
                i = str(i.encode("utf-8")).replace("<em>"," ").replace("</em>", " ")
                strText.append(i)

        #get the title and string it
        try:
            title = str(soup.title.text.replace('(Stanford Encyclopedia of Philosophy)', '')).rstrip()
            strText =    str(strText)[2:]
            x = re.sub(r'[^a-zA-Z]', " ", strText[0:])
        except:
            str(soup.title.text.encode('utf-8')).replace('(Stanford Encyclopedia of Philosophy)', '').rstrip()
            strText =    str(strText)[2:]
            x = re.sub(r'[^a-zA-Z]', " ", strText[0:])

        #get the authors and string them
        authors=[]
        removethese=['initial-scale','html' 'noarchive', title, '20', '19']
        try:
            thisauthor1 = soup.findAll('meta', attrs={'property': "citation_author"})
            authors.append(thisauthor1[0]['content'])
        except:
            thisauthor2 = soup.findAll('meta', attrs={'name': "citation_author"})
            authors.append(thisauthor2[0]['content'])

        #save the txt files from the scrap
        #these are inherently dirty text files that need to be cleaned.
        with open(workDir+os.sep+title.replace(" ", "")+".txt",'w') as f:
            f.write(title)
            f.write("\n")
            f.write(authors[0])
            f.write("\n")
            f.write(strText)

def getDocument(thetextfile=str):
    text=[]
    with open(workDir+os.sep+thetextfile, "rU") as f:
        for line in f:
            text.append(line)
    return text

def preprocess(textfile=list):
    """
    THIS IS STILL HORRIBLE BECAUSE IT CHECKS EVERY SINGLE WORD FOR ERRORS.
    """
    #put the text in a pandas.df and gl.SFrame for analysis
    df = pandas.DataFrame([['','','']], columns=['entry', 'authors', 'text'])
    df["entry"] = textfile[0].rstrip()
    df["authors"] = str(textfile[1]).replace("[","").replace("]","").lower().rstrip()
    df['text'] = graphlab.SArray([textfile[2]])

    #convert to graphlab SFrame
    df=graphlab.SFrame(df)

    # Do preprocessing of text
    df['text'] = graphlab.text_analytics.trim_rare_words(df['text'])

    #get and clean the word counts
    df['word_count']=graphlab.text_analytics.count_words(df['text'])
    df['word_count']=df['word_count'].dict_trim_by_keys(graphlab.text_analytics.stopwords(), True)
    wordcountdict={}
    for i in df['word_count'][0]:
        x=i
        if len(x)>3:
            if not x[0].isalpha():
                x = x[1:]
            if not x[-1].isalpha():
                x = x[0:-1]
            if not x[-1].isalpha():
                x = x[0:-1]
            if x.isalpha():
                if len(x)>2:
                    wordcountdict[x.rstrip()] = df['word_count'][0][i]
    df['word_count'] = graphlab.SArray([wordcountdict])
    del(wordcountdict)

    df_table = df[['word_count']].stack('word_count', new_column_name=['word', 'count'])
    #df['word']=df_table['word']
    #df['count']=df_table['count']

    df["top_word"] = None
    return df, df_table

def createCorpus(workDir=str, scrap=False):
    if scrap==True:
        scrapStanford(100)
    thetextfiles = [i for i in os.listdir(workDir+os.sep) if i.endswith(".txt")]
    textfile = getDocument(thetextfile=thetextfiles[0])
    df,df_table = preprocess(textfile=textfile)
    for i in thetextfiles[1:]:
        textfile = getDocument(thetextfile=i)
        df0,df_table0 = preprocess(textfile=textfile)
        df=df.append(df0)
    return df



def buildModels(df=graphlab.SFrame, show_sorted_word_count=False):

    df['tfidf'] = graphlab.text_analytics.tf_idf(df['word_count'])

    knn_model = graphlab.nearest_neighbors.create(df,features=['tfidf'],label='entry',distance='cosine')

    #CORPUS WORD COUNTS SORTED BY TF-IDF
    if show_sorted_word_count == True:
        d_table = df[['tfidf']].stack('tfidf', new_column_name=['word', 'count'])
        d_table.sort('count', ascending=False).print_rows(10)

    return df, knn_model

def showNetwork(sim_graph):
    # convert to NetworkX graph
    src = [sim_graph.edges['__src_id'][i][0:24].rstrip() for i in range(sim_graph.edges['__src_id'].shape[0])]
    dst = [sim_graph.edges['__dst_id'][i][0:24].rstrip() for i in range(sim_graph.edges['__dst_id'].shape[0])]
    #dis = [sim_graph.edges['distance'][i] for i in range(sim_graph.edges['distance'].shape[0])]
    #rank = [sim_graph.edges['rank'][i] for i in range(sim_graph.edges['distance'].shape[0])]
    connections = zip(src, dst)
    #del(src)
    #del(dst)
    g = nx.DiGraph()
    g.add_nodes_from([sim_graph.vertices['__id'][i][0:24].rstrip() for i in range(sim_graph.vertices['__id'].shape[0])])
    g.add_edges_from(connections)
    # Draw it and show on screen
    draw(g)
        #plt.show()

#if __name__ == "__main__":
#	df = createCorpus(workDir=workDir, scrap=False)
#	df, knn_model = buildModels(df=df, show_sorted_word_count=False)
#	sim_graph = knn_model.similarity_graph(k=3)
#	#sim_graph.show(vlabel='id', arrows=True);
#	showNetwork(sim_graph=sim_graph)
#	nearestNeighbors = knn_model.query(df, 'entry');
#	nearestNeighbors.print_rows(num_rows=180, num_columns=2)

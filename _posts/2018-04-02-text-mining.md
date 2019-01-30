---
layout: post
title: text-mining
date: 2018-04-02
author: dc
comments: true
---

I recently read <a href = "http://juanrloaiza.blogspot.de/2017/08/whos-most-mentioned-philosopher-in-sep.html"> a friend's post</a> on mining the Stanford Encyclopedia of Philosophy. I wanted to refresh my knowledge of KNN-Models and I spent some free time writing a Python script for scraping and analyzing random entries from the SEP. The image below shows the connections between the subset of the SEP that I mined. You can find the code for producing the image below, along with a table of the metrics.

I will summarize what I've learned some time else. But, in general, this looks like a promising method to uncover similarity within linguistic corpuses.

Note: <a href = "https://github.com/dnck/dnck.github.io/blob/master/assets/data/knn_on_sep.py">if you decide to run this code, the function preprocess() is still incredibly time consuming because I suck at sanitizing text.</a>

<div class="container-fluid">
	<div class="row">
		<div class = "col-md-6">
			<img src="{{site.url}}/assets/images/stanfordmine.png" class="img-fluid">
		</div>
	</div>
</div>
<br>

Figure 1. Directed Graph of the nearest neighbors for 35 articles. Scroll down to see another DA with 100 articles.


<div class="container-fluid">
	<div class="row">
		<div class = "col-md-6">
			<img src="{{site.url}}/assets/images/stanfordmine2.png" class="img-fluid">
		</div>
	</div>
</div>
<br>
Figure 2. Directed Graph of the nearest neighbors for 100 articles.



## Table 1. A Subset of articles in the SEP along with their nearest neighbors.

| query_label | reference_label | distance    
| ------------- |:-------------:|-------------:|
|         Social Minimum        |         Social Minimum        |		0
|         Social Minimum        |             Xunzi             |		0.886066449
|         Social Minimum        |      Political Obligation     |		0.893912798
|         Social Minimum        |      Natural Law Theories     |		0.8961885
|         Social Minimum        |      Transitional Justice     |		0.905404421
|         Ernst Cassirer        |         Ernst Cassirer        |		-2.22E-16
|         Ernst Cassirer        | Early Philosophical Interp... |		0.889126825
|         Ernst Cassirer        | 18th Century German Philos... |		0.89923616
|         Ernst Cassirer        | The Distinction Between In... |		0.910329509
|         Ernst Cassirer        | Ancient and Medieval Empir... |		0.917730917
| The Ethics and Rationality... | The Ethics and Rationality... |		2.22E-16
| The Ethics and Rationality... |      Political Obligation     |		0.907945565
| The Ethics and Rationality... |         Social Minimum        |		0.933382
| The Ethics and Rationality... |      Transitional Justice     |		0.956323289
| The Ethics and Rationality... |        Informed Consent       |		0.968030454
|        Turing Machines        |        Turing Machines        |		0
|        Turing Machines        |         Social Minimum        |		0.979068005
|        Turing Machines        |            Dualism            |		0.985629512
|        Turing Machines        | Quantum Theory: von Neuman... |		0.986195312
|        Turing Machines        |          Dialetheism          |		0.987783281
|        Value Pluralism        |        Value Pluralism        |		-2.22E-16
|        Value Pluralism        |    Charles Leslie Stevenson   |		0.948372664
|        Value Pluralism        |         Social Minimum        |		0.949139019
|        Value Pluralism        |             Xunzi             |		0.951399153
|        Value Pluralism        |    The Cambridge Platonists   |		0.955464692
| The Distinction Between In... | The Distinction Between In... |		-2.22E-16
| The Distinction Between In... |          Behaviorism          |		0.894853202
| The Distinction Between In... |         Ernst Cassirer        |		0.910329509
| The Distinction Between In... | Ancient and Medieval Empir... |		0.932283262
| The Distinction Between In... |            Dualism            |		0.948019733
|          Sovereignty          |          Sovereignty          |		0
|          Sovereignty          |      Transitional Justice     |		0.870650203
|          Sovereignty          |      Political Obligation     |		0.886332241
|          Sovereignty          |      Natural Law Theories     |		0.905935222
|          Sovereignty          |    The Cambridge Platonists   |		0.948555235
|             Elias             |             Elias             |		1.11E-16
|             Elias             | 18th Century German Philos... |		0.953050014
|             Elias             | Ancient and Medieval Empir... |		0.959690798
|             Elias             |        Albert the Great       |		0.968017414
|             Elias             |           Gersonides          |		0.96944534
|           Gersonides          |           Gersonides          |		2.22E-16
|           Gersonides          | Ancient and Medieval Empir... |		0.885827927
|           Gersonides          |        Albert the Great       |		0.890411638
|           Gersonides          |    The Cambridge Platonists   |		0.935804286
|           Gersonides          |            Dualism            |		0.944821267
| Quantum Theory: von Neuman... | Quantum Theory: von Neuman... |		0
| Quantum Theory: von Neuman... | Singularities and Black Holes |		0.846729482
| Quantum Theory: von Neuman... | Early Philosophical Interp... |		0.891706686
| Quantum Theory: von Neuman... |         Ernst Cassirer        |		0.950745344
| Quantum Theory: von Neuman... |        Benjamin Peirce        |		0.956408695
| Singularities and Black Holes | Singularities and Black Holes |		0
| Singularities and Black Holes | Early Philosophical Interp... |		0.762704457
| Singularities and Black Holes | Quantum Theory: von Neuman... |		0.846729482
| Singularities and Black Holes |            Dualism            |		0.931282974
| Singularities and Black Holes |           Gersonides          |		0.969017047
|      Transitional Justice     |      Transitional Justice     |		0
|      Transitional Justice     |      Natural Law Theories     |		0.866667716
|      Transitional Justice     |          Sovereignty          |		0.870650203
|      Transitional Justice     |      Political Obligation     |		0.887859582
|      Transitional Justice     |         Social Minimum        |		0.905404421
|            Dualism            |            Dualism            |		0
|            Dualism            | Ancient and Medieval Empir... |		0.882338477
|            Dualism            | Early Philosophical Interp... |		0.886307301
|            Dualism            |           Categories          |		0.896618063
|            Dualism            |   The Problem of Perception   |		0.903920633
|            Fideism            |            Fideism            |		0
|            Fideism            |          Pierre Bayle         |		0.815617917
|            Fideism            |    The Cambridge Platonists   |		0.918675784
|            Fideism            |             Xunzi             |		0.934754893
|            Fideism            | Ancient and Medieval Empir... |		0.935299253
| Ancient and Medieval Empir... | Ancient and Medieval Empir... |		0
| Ancient and Medieval Empir... |        Albert the Great       |		0.881518682
| Ancient and Medieval Empir... |            Dualism            |		0.882338477
| Ancient and Medieval Empir... |           Gersonides          |		0.885827927
| Ancient and Medieval Empir... |       Medieval Mereology      |		0.909565352
|           Categories          |           Categories          |		-2.22E-16
|           Categories          |            Dualism            |		0.896618063
|           Categories          |             Xunzi             |		0.928751967
|           Categories          | Ancient and Medieval Empir... |		0.939257328
|           Categories          |      Episteme and Techne      |		0.939883182
|       Philosophy of Film      |       Philosophy of Film      |		-2.22E-16
|       Philosophy of Film      |         Social Minimum        |		0.967254136
|       Philosophy of Film      | 18th Century German Philos... |		0.97486587
|       Philosophy of Film      |             Xunzi             |		0.975322957
|       Philosophy of Film      |         Ernst Cassirer        |		0.978168325
|       Charles Hartshorne      |       Charles Hartshorne      |		0
|       Charles Hartshorne      |            Fideism            |		0.945839597
|       Charles Hartshorne      | Ancient and Medieval Empir... |		0.949117033
|       Charles Hartshorne      |    The Cambridge Platonists   |		0.955241106
|       Charles Hartshorne      |           Gersonides          |		0.959512044
|        Albert the Great       |        Albert the Great       |		1.11E-16
|        Albert the Great       | Ancient and Medieval Empir... |		0.881518682
|        Albert the Great       |           Gersonides          |		0.890411638
|        Albert the Great       |            Dualism            |		0.935281893
|        Albert the Great       |    The Cambridge Platonists   |		0.935845122
|   The Problem of Perception   |   The Problem of Perception   |		0
|   The Problem of Perception   |            Dualism            |		0.903920633
|   The Problem of Perception   |          Behaviorism          |		0.951424804
|   The Problem of Perception   |         Ernst Cassirer        |		0.957079125
|   The Problem of Perception   | Ancient and Medieval Empir... |		0.96042406
|      Natural Law Theories     |      Natural Law Theories     |		0
|      Natural Law Theories     |      Political Obligation     |		0.854791875
|      Natural Law Theories     |      Transitional Justice     |		0.866667716
|      Natural Law Theories     |         Social Minimum        |		0.8961885
|      Natural Law Theories     |          Sovereignty          |		0.905935222
|             Xunzi             |             Xunzi             |		1.11E-16
|             Xunzi             |         Social Minimum        |		0.886066449
|             Xunzi             |      Natural Law Theories     |		0.917299131
|             Xunzi             |      Episteme and Techne      |		0.923708702
|             Xunzi             | 18th Century German Philos... |		0.924618655
|      Episteme and Techne      |      Episteme and Techne      |		2.22E-16
|      Episteme and Techne      |       Medieval Mereology      |		0.878755061
|      Episteme and Techne      | Ancient and Medieval Empir... |		0.911243664
|      Episteme and Techne      |             Xunzi             |		0.923708702
|      Episteme and Techne      |         Ernst Cassirer        |		0.928558842
|        Informed Consent       |        Informed Consent       |		0
|        Informed Consent       |    Decision-Making Capacity   |		0.835186086
|        Informed Consent       |         Social Minimum        |		0.943259231
|        Informed Consent       |      Political Obligation     |		0.949249943
|        Informed Consent       |      Transitional Justice     |		0.950165726
|    Charles Leslie Stevenson   |    Charles Leslie Stevenson   |		-2.22E-16
|    Charles Leslie Stevenson   |             Xunzi             |		0.933961739
|    Charles Leslie Stevenson   |            Dualism            |		0.936582461
|    Charles Leslie Stevenson   |         Social Minimum        |		0.942144503
|    Charles Leslie Stevenson   |        Value Pluralism        |		0.948372664
|        Benjamin Peirce        |        Benjamin Peirce        |		0
|        Benjamin Peirce        | 18th Century German Philos... |		0.929544331
|        Benjamin Peirce        |         Ernst Cassirer        |		0.946998499
|        Benjamin Peirce        | Quantum Theory: von Neuman... |		0.956408695
|        Benjamin Peirce        | Early Philosophical Interp... |		0.965199743
|          Pierre Bayle         |          Pierre Bayle         |		0
|          Pierre Bayle         |            Fideism            |		0.815617917
|          Pierre Bayle         | 18th Century German Philos... |		0.924485785
|          Pierre Bayle         |    The Cambridge Platonists   |		0.933286716
|          Pierre Bayle         |            Dualism            |		0.93427225
|       Medieval Mereology      |       Medieval Mereology      |		0
|       Medieval Mereology      |      Episteme and Techne      |		0.878755061
|       Medieval Mereology      | Ancient and Medieval Empir... |		0.909565352
|       Medieval Mereology      |        Albert the Great       |		0.936429009
|       Medieval Mereology      |           Categories          |		0.958510668
|    Decision-Making Capacity   |    Decision-Making Capacity   |		2.22E-16
|    Decision-Making Capacity   |        Informed Consent       |		0.835186086
|    Decision-Making Capacity   |      Natural Law Theories     |		0.921007459
|    Decision-Making Capacity   |             Xunzi             |		0.948070092
|    Decision-Making Capacity   |      Political Obligation     |		0.957150923
|      Political Obligation     |      Political Obligation     |		0
|      Political Obligation     |      Natural Law Theories     |		0.854791875
|      Political Obligation     |          Sovereignty          |		0.886332241
|      Political Obligation     |      Transitional Justice     |		0.887859582
|      Political Obligation     |         Social Minimum        |		0.893912798
|    The Cambridge Platonists   |    The Cambridge Platonists   |		0
|    The Cambridge Platonists   |            Fideism            |		0.918675784
|    The Cambridge Platonists   | Ancient and Medieval Empir... |		0.926492079
|    The Cambridge Platonists   |          Pierre Bayle         |		0.933286716
|    The Cambridge Platonists   |           Gersonides          |		0.935804286
|          Behaviorism          |          Behaviorism          |		1.11E-16
|          Behaviorism          | The Distinction Between In... |		0.894853202
|          Behaviorism          |            Dualism            |		0.939802485
|          Behaviorism          |             Xunzi             |		0.946861839
|          Behaviorism          |   The Problem of Perception   |		0.951424804
| Early Philosophical Interp... | Early Philosophical Interp... |		0
| Early Philosophical Interp... | Singularities and Black Holes |		0.762704457
| Early Philosophical Interp... |            Dualism            |		0.886307301
| Early Philosophical Interp... |         Ernst Cassirer        |		0.889126825
| Early Philosophical Interp... | Quantum Theory: von Neuman... |		0.891706686
| 18th Century German Philos... | 18th Century German Philos... |		1.11E-16
| 18th Century German Philos... |         Ernst Cassirer        |		0.89923616
| 18th Century German Philos... |          Pierre Bayle         |		0.924485785
| 18th Century German Philos... |             Xunzi             |		0.924618655
| 18th Century German Philos... |        Benjamin Peirce        |		0.929544331
|      Logical Consequence      |      Logical Consequence      |		1.11E-16
|      Logical Consequence      |          Dialetheism          |		0.893238821
|      Logical Consequence      | Ancient and Medieval Empir... |		0.954412798
|      Logical Consequence      |         Ernst Cassirer        |		0.96222519
|      Logical Consequence      |           Categories          |		0.962310752
|          Dialetheism          |          Dialetheism          |		0
|          Dialetheism          |      Logical Consequence      |		0.893238821
|          Dialetheism          |      Episteme and Techne      |		0.930745099
|          Dialetheism          |         Ernst Cassirer        |		0.936979938
|          Dialetheism          |            Dualism            |		0.938811062

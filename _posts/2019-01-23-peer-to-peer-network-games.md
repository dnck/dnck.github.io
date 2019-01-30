---
layout: post
title: P2P network games
author: dc
date: 2019-01-22
comments: true
analytics: true
keywords:  peer-to-peer networks
description: an imaginary game for peer to peer networks   
category: peer-to-peer networks
---

Games

Let us think about a special two player game in which if one player successfully cheats, then the other player losses the value gained by the former. On the other hand, if the cheater is detected, then the cheater losses a staked value to the non-cheater. Finally, in this game, both players will gain some value from cooperating and not cheating.

Table 1. shows the combination of player inputs and the outcomes of the game when the staked values are standardized to 1.

|   |   |   |   |   |   |   |   |
|---|---|---|---|---|---|---|---|
| Player A Cheats |   1  | 1  | 0  | 0  | 0  | 1  | 1  |
| Player B Cheats |   0  | 1  | 1  | 0  | 1  | 0  | 1  |
| Player A is successful |   1  | 1  | 0  | 0  | 0  | 0  | 0  |
| Player B is successful |  0  | 1  | 1  | 0  | 0  | 0  | 0  |
| Player A Outcome/Gain | 1  | 0  | 0  | 1  |   |   | 0  |
| Player B Outcome/Gain  | 0  | 0  | 1  | 1  |   |   | 0  |
| Player A Outcome/Loss | 0  | -1  | -1  | 0   | 0  | -2  | -2  |
| Player B Outcome/Loss | -1  | -1  | 0  | 0  | -2  | 0  | -2  |

The way to read Table 1. is to start in the left most column. Reading down row-wise, we get the sequence, [1,0,1,0,1,0,0,-1]. We interpret this as follows: if Player A cheats, is successful, and Player B does not cheat and is (obviously) not successful at cheating, then, as per the rules of the game, Player A gains a value of +1, but Player B loses the value of 1. A similar logic can be applied to the other columns in the table.

What can be seen from Table 1. is that the biggest gains for both players come when they do not cheat, but their partners are exposed as cheating. After that, the gains for both players are the same whether they cheat or not. Thus, the best option for both Players is to not cheat, since their biggest win comes from not cheating, their biggest loss from failed attempts to cheat, and additionally, the benefit they stand to gain from cheating and not being discovered is as large as the benefit from not cheating when their partner also does not cheat. It seems then, that this sort of game encourages cooperation among the players.

Now, take the same game logic and apply it to the hosts in a peer-to-peer network. The hosts in the network will implement a permissionless many-read-once-write database consisting of timestamped records of the existence of a digital asset. What we would like to offer all participants in the network is the guarantee that the timestamp of a transaction is equal to the first timestamp on the transaction when its read from a local instantiation of the shared database, and plus, the information can be advertised by the autonomous agent that originally broadcasted the transaction. We can do this by building the above game logic into our software design.


A proposal for a new game


Suppose that when a node wishes to broadcast a transaction, they must receive a timestamp on the transaction from one of their neighbor nodes.
To get a timestamp from a neighbor, a sender signs a data file that contains at least, a unique transaction address, a transaction timestamp, a broadcast interval, two empty variables to be explained later, the sender id (certified host:port), and the recipient id (certified host:port). Neighbors are incentivized to accept these timestamp requests because all nodes are continually adding transactions to their local databases. We assume that some of these transactions issued by nodes are intended for a shared public database, which acts as the network's global ledger, and that the only way into this ledger is through a network consensus approved timestamp.

Now, an important rule for the game is that when a node sends another node a timestamp request, the recipient should do two things. First, the recipient should incorporate the transaction into his local DAG in the aforementioned way. Second, the recipient should return to the sender a proof of the transaction's incorporation into the recipient's local DAG. Thus, sending a transaction from one node to another node in this game constitutes a request for contract. We explain these steps below.

First, when a neighbor node adds the transaction of their peer to their own database, they will always hash the new transaction with two transactions in their database that are not yet hashed by any other transactions. In other words, each new transaction is linked to k older transactions (where k can be chosen by the network consensus). Occasionally, these transactions will be value transferring transactions, but mostly (at least, at the start of the network), the transactions will not hold any value. Thus, when the receiver accepts a sender's transaction timestamp request, the receiver will follow these steps for incorporating the new transaction into her local database.

Second, what the receiver should return to the sender is a shared secret. The shared secret will contains the sender's signed copy of the original transaction request, and, additionally, the transactions location in the recipient's database and its path.

An important part of the shared secret is the transactions location in the recipient's database and its path information. Recall that the transactions path is generated by the concatenation of the sender's transaction at time t with information from the receiver's database state at time t + 1. Each nodes database state for all times > 0 will be an instance of a Directed Acyclic Graph. Their local DAGs will consist of many transactions approving other transactions. Thus, what the recipient returns to the sender is the path information of the sender's new transaction in the recipients database.

The returned shared secret will also be associated with an interval of time called, the broadcast interval. During this interval, the sender is permitted to request from the receiver that the receiver broadcasts the secret to which the sender can testify knowledge. After the interval has expired, the receiver can likewise broadcast the revealed secret.

After the sender and receiver know of the transaction and the secret, both are now free to broadcast the transaction to other nodes. During this broadcast interval, an option is selected that tells the new receivers that the transaction is already in contract.

Here are the rules of the broadcast interval.

If the sender's reveal request comes within the agreed upon broadcast interval, but the receiver does not broadcast the secret, then the sender can reveal the contents of the secret. In this case, the sender's revelation of the secret should show that the receiver has cheated, and the receiver should be penalized by 2, and the sender rewarded by 2.

On the other hand, if the sender does not request from the receiver the reveal of the secret within the agreed upon interval of time, then the receiver can broadcast the revelation of the secret, and the sender should be penalized by 2, and the receiver rewarded by 2.

From this it is clear that both sender and receiver should execute their respective broadcast interval functions in due time. In a case of unintentional cheating, e.g. crash failures, the network can accept crash reports from the failed autonomous agent, and make tribunal decisions about liabilities for damages.

Recall now that the goal of this game is to provide consistent timestamps for many reads on the same transaction accessed from any local ledger, and that these timestamps are equal to the first recorded time by the sender.

To make this system work in providing for our goal, we should use a third type of node in the game that aggregates all value bearing transactions and their metadata, and puts this information into a storage container also organized as an instance of a Directed Acyclic Graph.

Questions and Discussion

There are many questions that need to be answered about this proposal before evaluating it. Here are just a few:

First, what prevents the nodes from cheating?

Second, what are the rewards for not cheating?

Third, how are punishments implemented?

Fourth, how can nodes cheat?

Fifth, what's the purpose of having the nodes continually add transactions?

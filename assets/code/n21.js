<script src="https://d3js.org/d3.v5.min.js"></script>
<script>
var nodes = [
  {id: "0", group: 0, label: "N0", level: 2 },
  {id: "1", group: 0, label: "N1", level: 1},
  {id: "2", group: 0, label: "N2", level: 2 },
  {id: "3", group: 0, label: "*N3*", level: 1 },
  {id: "4", group: 0, label: "N4", level: 2 },
  {id: "5", group: 0, label: "N5", level: 2 },
  {id: "6", group: 0, label: "N6", level: 2 },
  {id: "7", group: 0, label: "N7", level: 2 },
  {id: "8", group: 0, label: "N8", level: 2 },
  {id: "9", group: 0, label: "N9", level: 2 },
  {id: "10", group: 0, label: "N10", level: 1},
  {id: "11", group: 0, label: "N11", level: 2},
  {id: "12", group: 0, label: "N12", level: 2},
  {id: "13", group: 0, label: "N13", level: 2},
  {id: "14", group: 0, label: "N14", level: 1},
  {id: "15", group: 0, label: "N15", level: 2},
  {id: "16", group: 0, label: "N16", level: 2},
  {id: "17", group: 0, label: "N17", level: 2},
  {id: "18", group: 0, label: "N18", level: 2},
  {id: "19", group: 0, label: "N19", level: 2},
  {id: "20", group: 0, label: "N20", level: 2}
]

var links = [
		{target: "10", source: "0" , strength: 0.1},
		{target: "6", source: "0" , strength: 0.1},
		{target: "7", source: "0" , strength: 0.1},
		{target: "2", source: "0" , strength: 0.1},
		{target: "12", source: "0" , strength: 0.1},
		{target: "17", source: "0" , strength: 0.1},
		{target: "9", source: "0" , strength: 0.1},
		{target: "11", source: "0" , strength: 0.1},
		{target: "3", source: "1" , strength: 0.1},
		{target: "14", source: "1" , strength: 0.1},
		{target: "8", source: "1" , strength: 0.1},
		{target: "16", source: "1" , strength: 0.1},
		{target: "18", source: "1" , strength: 0.1},
		{target: "13", source: "1" , strength: 0.1},
		{target: "9", source: "1" , strength: 0.1},
		{target: "20", source: "1" , strength: 0.1},
		{target: "12", source: "1" , strength: 0.1},
		{target: "15", source: "1" , strength: 0.1},
		{target: "17", source: "2" , strength: 0.1},
		{target: "0", source: "2" , strength: 0.1},
		{target: "11", source: "2" , strength: 0.1},
		{target: "10", source: "2" , strength: 0.1},
		{target: "5", source: "2" , strength: 0.1},
		{target: "7", source: "2" , strength: 0.1},
		{target: "15", source: "2" , strength: 0.1},
		{target: "14", source: "3" , strength: 0.1},
		{target: "1", source: "3" , strength: 0.1},
		{target: "10", source: "3" , strength: 0.1},
		{target: "20", source: "4" , strength: 0.1},
		{target: "13", source: "4" , strength: 0.1},
		{target: "15", source: "4" , strength: 0.1},
		{target: "16", source: "4" , strength: 0.1},
		{target: "5", source: "4" , strength: 0.1},
		{target: "19", source: "4" , strength: 0.1},
		{target: "11", source: "5" , strength: 0.1},
		{target: "15", source: "5" , strength: 0.1},
		{target: "17", source: "5" , strength: 0.1},
		{target: "19", source: "5" , strength: 0.1},
		{target: "2", source: "5" , strength: 0.1},
		{target: "4", source: "5" , strength: 0.1},
		{target: "6", source: "5" , strength: 0.1},
		{target: "20", source: "5" , strength: 0.1},
		{target: "18", source: "5" , strength: 0.1},
		{target: "0", source: "6" , strength: 0.1},
		{target: "10", source: "6" , strength: 0.1},
		{target: "17", source: "6" , strength: 0.1},
		{target: "7", source: "6" , strength: 0.1},
		{target: "11", source: "6" , strength: 0.1},
		{target: "12", source: "6" , strength: 0.1},
		{target: "5", source: "6" , strength: 0.1},
		{target: "19", source: "6" , strength: 0.1},
		{target: "12", source: "7" , strength: 0.1},
		{target: "10", source: "7" , strength: 0.1},
		{target: "9", source: "7" , strength: 0.1},
		{target: "0", source: "7" , strength: 0.1},
		{target: "18", source: "7" , strength: 0.1},
		{target: "6", source: "7" , strength: 0.1},
		{target: "8", source: "7" , strength: 0.1},
		{target: "2", source: "7" , strength: 0.1},
		{target: "1", source: "8" , strength: 0.1},
		{target: "18", source: "8" , strength: 0.1},
		{target: "14", source: "8" , strength: 0.1},
		{target: "9", source: "8" , strength: 0.1},
		{target: "16", source: "8" , strength: 0.1},
		{target: "12", source: "8" , strength: 0.1},
		{target: "13", source: "8" , strength: 0.1},
		{target: "7", source: "8" , strength: 0.1},
		{target: "8", source: "9" , strength: 0.1},
		{target: "7", source: "9" , strength: 0.1},
		{target: "1", source: "9" , strength: 0.1},
		{target: "14", source: "9" , strength: 0.1},
		{target: "0", source: "9" , strength: 0.1},
		{target: "13", source: "9" , strength: 0.1},
		{target: "19", source: "9" , strength: 0.1},
		{target: "11", source: "9" , strength: 0.1},
		{target: "3", source: "10" , strength: 0.1},
		{target: "7", source: "10" , strength: 0.1},
		{target: "0", source: "10" , strength: 0.1},
		{target: "12", source: "10" , strength: 0.1},
		{target: "6", source: "10" , strength: 0.1},
		{target: "2", source: "10" , strength: 0.1},
		{target: "18", source: "10" , strength: 0.1},
		{target: "16", source: "10" , strength: 0.1},
		{target: "17", source: "11" , strength: 0.1},
		{target: "5", source: "11" , strength: 0.1},
		{target: "2", source: "11" , strength: 0.1},
		{target: "6", source: "11" , strength: 0.1},
		{target: "19", source: "11" , strength: 0.1},
		{target: "0", source: "11" , strength: 0.1},
		{target: "18", source: "11" , strength: 0.1},
		{target: "9", source: "11" , strength: 0.1},
		{target: "16", source: "11" , strength: 0.1},
		{target: "7", source: "12" , strength: 0.1},
		{target: "18", source: "12" , strength: 0.1},
		{target: "10", source: "12" , strength: 0.1},
		{target: "8", source: "12" , strength: 0.1},
		{target: "0", source: "12" , strength: 0.1},
		{target: "1", source: "12" , strength: 0.1},
		{target: "6", source: "12" , strength: 0.1},
		{target: "20", source: "13" , strength: 0.1},
		{target: "16", source: "13" , strength: 0.1},
		{target: "4", source: "13" , strength: 0.1},
		{target: "14", source: "13" , strength: 0.1},
		{target: "1", source: "13" , strength: 0.1},
		{target: "19", source: "13" , strength: 0.1},
		{target: "8", source: "13" , strength: 0.1},
		{target: "15", source: "13" , strength: 0.1},
		{target: "9", source: "13" , strength: 0.1},
		{target: "3", source: "14" , strength: 0.1},
		{target: "16", source: "14" , strength: 0.1},
		{target: "1", source: "14" , strength: 0.1},
		{target: "13", source: "14" , strength: 0.1},
		{target: "8", source: "14" , strength: 0.1},
		{target: "20", source: "14" , strength: 0.1},
		{target: "18", source: "14" , strength: 0.1},
		{target: "9", source: "14" , strength: 0.1},
		{target: "5", source: "15" , strength: 0.1},
		{target: "4", source: "15" , strength: 0.1},
		{target: "17", source: "15" , strength: 0.1},
		{target: "20", source: "15" , strength: 0.1},
		{target: "2", source: "15" , strength: 0.1},
		{target: "13", source: "15" , strength: 0.1},
		{target: "1", source: "15" , strength: 0.1},
		{target: "13", source: "16" , strength: 0.1},
		{target: "14", source: "16" , strength: 0.1},
		{target: "20", source: "16" , strength: 0.1},
		{target: "1", source: "16" , strength: 0.1},
		{target: "4", source: "16" , strength: 0.1},
		{target: "8", source: "16" , strength: 0.1},
		{target: "18", source: "16" , strength: 0.1},
		{target: "19", source: "16" , strength: 0.1},
		{target: "10", source: "16" , strength: 0.1},
		{target: "11", source: "16" , strength: 0.1},
		{target: "2", source: "17" , strength: 0.1},
		{target: "11", source: "17" , strength: 0.1},
		{target: "6", source: "17" , strength: 0.1},
		{target: "5", source: "17" , strength: 0.1},
		{target: "0", source: "17" , strength: 0.1},
		{target: "15", source: "17" , strength: 0.1},
		{target: "19", source: "17" , strength: 0.1},
		{target: "8", source: "18" , strength: 0.1},
		{target: "1", source: "18" , strength: 0.1},
		{target: "12", source: "18" , strength: 0.1},
		{target: "14", source: "18" , strength: 0.1},
		{target: "7", source: "18" , strength: 0.1},
		{target: "16", source: "18" , strength: 0.1},
		{target: "10", source: "18" , strength: 0.1},
		{target: "5", source: "18" , strength: 0.1},
		{target: "11", source: "18" , strength: 0.1},
		{target: "5", source: "19" , strength: 0.1},
		{target: "20", source: "19" , strength: 0.1},
		{target: "11", source: "19" , strength: 0.1},
		{target: "13", source: "19" , strength: 0.1},
		{target: "17", source: "19" , strength: 0.1},
		{target: "16", source: "19" , strength: 0.1},
		{target: "6", source: "19" , strength: 0.1},
		{target: "9", source: "19" , strength: 0.1},
		{target: "4", source: "19" , strength: 0.1},
		{target: "4", source: "20" , strength: 0.1},
		{target: "13", source: "20" , strength: 0.1},
		{target: "16", source: "20" , strength: 0.1},
		{target: "19", source: "20" , strength: 0.1},
		{target: "14", source: "20" , strength: 0.1},
		{target: "15", source: "20" , strength: 0.1},
		{target: "1", source: "20" , strength: 0.1},
		{target: "5", source: "20" , strength: 0.1}
]

var endpoint_averages = {
    "0": {
        "1": "518.7",
        "10": "519.4",
        "11": "508.8",
        "12": "492.6",
        "13": "516.5",
        "14": "503.4",
        "15": "515.2",
        "16": "502.4",
        "17": "520.4",
        "18": "506.6",
        "19": "496.7",
        "2": "492.6",
        "20": "505.6",
        "3": "63.15",
        "4": "522.4",
        "5": "497.6",
        "6": "517.2",
        "7": "504.9",
        "8": "499.0",
        "9": "501.4"
    },
    "1": {
        "0": "493.1",
        "10": "520.9",
        "11": "510.5",
        "12": "492.6",
        "13": "514.5",
        "14": "501.2",
        "15": "513.7",
        "16": "500.7",
        "17": "522.5",
        "18": "505.1",
        "19": "496.6",
        "2": "495.0",
        "20": "502.9",
        "3": "58.92",
        "4": "520.1",
        "5": "497.9",
        "6": "519.3",
        "7": "506.3",
        "8": "497.3",
        "9": "501.4"
    },
    "10": {
        "0": "491.3",
        "1": "518.6",
        "11": "510.1",
        "12": "492.5",
        "13": "516.4",
        "14": "502.7",
        "15": "515.5",
        "16": "500.9",
        "17": "521.7",
        "18": "505.3",
        "19": "497.1",
        "2": "492.7",
        "20": "504.8",
        "3": "62.47",
        "4": "521.0",
        "5": "497.9",
        "6": "517.1",
        "7": "504.9",
        "8": "498.9",
        "9": "502.8"
    },
    "11": {
        "0": "491.5",
        "1": "518.6",
        "10": "520.8",
        "12": "494.4",
        "13": "516.2",
        "14": "502.7",
        "15": "514.9",
        "16": "501.0",
        "17": "520.3",
        "18": "505.3",
        "19": "495.2",
        "2": "492.7",
        "20": "504.4",
        "3": "63.15",
        "4": "520.7",
        "5": "496.4",
        "6": "517.3",
        "7": "506.5",
        "8": "499.1",
        "9": "501.6"
    },
    "12": {
        "0": "491.4",
        "1": "517.2",
        "10": "519.5",
        "11": "510.3",
        "13": "516.2",
        "14": "502.8",
        "15": "515.8",
        "16": "502.2",
        "17": "522.1",
        "18": "505.2",
        "19": "497.4",
        "2": "494.2",
        "20": "505.0",
        "3": "61.81",
        "4": "522.0",
        "5": "498.2",
        "6": "517.4",
        "7": "504.9",
        "8": "497.5",
        "9": "502.8"
    },
    "13": {
        "0": "493.4",
        "1": "516.9",
        "10": "521.2",
        "11": "510.3",
        "12": "494.1",
        "14": "501.1",
        "15": "513.6",
        "16": "500.5",
        "17": "522.3",
        "18": "506.3",
        "19": "495.2",
        "2": "494.8",
        "20": "502.5",
        "3": "60.74",
        "4": "518.5",
        "5": "497.8",
        "6": "520.0",
        "7": "506.7",
        "8": "497.4",
        "9": "501.4"
    },
    "14": {
        "0": "493.3",
        "1": "516.8",
        "10": "521.0",
        "11": "510.2",
        "12": "493.8",
        "13": "514.2",
        "15": "515.0",
        "16": "500.4",
        "17": "523.2",
        "18": "505.1",
        "19": "496.5",
        "2": "495.5",
        "20": "502.7",
        "3": "58.49",
        "4": "519.8",
        "5": "498.1",
        "6": "519.9",
        "7": "506.4",
        "8": "497.3",
        "9": "501.4"
    },
    "15": {
        "0": "493.2",
        "1": "517.1",
        "10": "521.5",
        "11": "510.1",
        "12": "494.8",
        "13": "514.6",
        "14": "502.7",
        "16": "502.1",
        "17": "520.4",
        "18": "506.7",
        "19": "496.5",
        "2": "492.8",
        "20": "502.8",
        "3": "61.75",
        "4": "518.6",
        "5": "496.4",
        "6": "519.0",
        "7": "507.1",
        "8": "499.2",
        "9": "503.2"
    },
    "16": {
        "0": "493.2",
        "1": "516.9",
        "10": "519.6",
        "11": "508.9",
        "12": "494.0",
        "13": "514.4",
        "14": "501.0",
        "15": "515.0",
        "17": "522.4",
        "18": "505.3",
        "19": "495.5",
        "2": "494.4",
        "20": "502.6",
        "3": "60.89",
        "4": "519.5",
        "5": "497.9",
        "6": "519.0",
        "7": "506.6",
        "8": "497.5",
        "9": "502.7"
    },
    "17": {
        "0": "491.5",
        "1": "519.1",
        "10": "520.9",
        "11": "508.9",
        "12": "494.4",
        "13": "516.7",
        "14": "504.2",
        "15": "513.7",
        "16": "502.8",
        "18": "507.0",
        "19": "495.5",
        "2": "492.6",
        "20": "504.5",
        "3": "63.97",
        "4": "520.7",
        "5": "496.5",
        "6": "517.3",
        "7": "506.6",
        "8": "500.4",
        "9": "503.2"
    },
    "18": {
        "0": "492.8",
        "1": "517.0",
        "10": "519.5",
        "11": "508.9",
        "12": "492.6",
        "13": "515.8",
        "14": "501.3",
        "15": "515.3",
        "16": "500.8",
        "17": "522.1",
        "19": "496.6",
        "2": "494.1",
        "20": "504.2",
        "3": "60.96",
        "4": "520.8",
        "5": "496.6",
        "6": "518.6",
        "7": "505.1",
        "8": "497.4",
        "9": "502.7"
    },
    "19": {
        "0": "492.5",
        "1": "518.3",
        "10": "521.4",
        "11": "508.7",
        "12": "494.7",
        "13": "513.9",
        "14": "502.3",
        "15": "513.9",
        "16": "500.2",
        "17": "519.9",
        "18": "506.4",
        "2": "493.9",
        "20": "501.8",
        "3": "63.44",
        "4": "518.9",
        "5": "495.7",
        "6": "516.3",
        "7": "506.6",
        "8": "499.0",
        "9": "500.5"
    },
    "2": {
        "0": "491.4",
        "1": "519.4",
        "10": "519.5",
        "11": "508.9",
        "12": "493.9",
        "13": "516.6",
        "14": "504.1",
        "15": "513.6",
        "16": "502.5",
        "17": "520.3",
        "18": "506.8",
        "19": "496.4",
        "20": "504.5",
        "3": "63.94",
        "4": "520.4",
        "5": "496.6",
        "6": "518.5",
        "7": "505.1",
        "8": "499.7",
        "9": "502.9"
    },
    "20": {
        "0": "494.3",
        "1": "517.0",
        "10": "521.4",
        "11": "510.3",
        "12": "494.6",
        "13": "514.3",
        "14": "501.2",
        "15": "513.6",
        "16": "500.6",
        "17": "522.0",
        "18": "506.5",
        "19": "495.2",
        "2": "494.5",
        "3": "60.74",
        "4": "518.2",
        "5": "496.5",
        "6": "519.4",
        "7": "507.8",
        "8": "498.7",
        "9": "502.7"
    },
    "3": {
        "0": "564.8",
        "1": "589.4",
        "10": "592.1",
        "11": "583.2",
        "12": "568.2",
        "13": "588.1",
        "14": "575.3",
        "15": "586.1",
        "16": "575.3",
        "17": "596.8",
        "18": "580.1",
        "19": "570.6",
        "2": "567.7",
        "20": "575.3",
        "4": "592.3",
        "5": "569.3",
        "6": "592.3",
        "7": "579.7",
        "8": "572.5",
        "9": "576.7"
    },
    "4": {
        "0": "494.1",
        "1": "518.3",
        "10": "521.3",
        "11": "510.1",
        "12": "495.3",
        "13": "514.1",
        "14": "502.2",
        "15": "513.3",
        "16": "500.5",
        "17": "521.7",
        "18": "506.7",
        "19": "495.4",
        "2": "494.1",
        "20": "502.4",
        "3": "62.26",
        "5": "496.3",
        "6": "519.1",
        "7": "507.8",
        "8": "498.8",
        "9": "503.0"
    },
    "5": {
        "0": "492.9",
        "1": "518.6",
        "10": "521.2",
        "11": "508.7",
        "12": "494.6",
        "13": "516.0",
        "14": "503.0",
        "15": "513.5",
        "16": "502.3",
        "17": "520.5",
        "18": "505.4",
        "19": "494.8",
        "2": "492.9",
        "20": "503.0",
        "3": "63.22",
        "4": "518.6",
        "6": "517.5",
        "7": "506.8",
        "8": "499.6",
        "9": "503.5"
    },
    "6": {
        "0": "487.2",
        "1": "515.1",
        "10": "515.3",
        "11": "504.6",
        "12": "488.5",
        "13": "513.2",
        "14": "500.2",
        "15": "511.0",
        "16": "498.3",
        "17": "516.2",
        "18": "502.5",
        "19": "491.3",
        "2": "489.8",
        "20": "500.6",
        "3": "59.72",
        "4": "517.7",
        "5": "492.4",
        "7": "500.8",
        "8": "495.0",
        "9": "498.7"
    },
    "7": {
        "0": "491.0",
        "1": "518.2",
        "10": "519.0",
        "11": "510.0",
        "12": "492.2",
        "13": "515.9",
        "14": "502.4",
        "15": "515.2",
        "16": "501.9",
        "17": "521.4",
        "18": "505.0",
        "19": "496.8",
        "2": "492.5",
        "20": "505.2",
        "3": "62.31",
        "4": "521.4",
        "5": "497.6",
        "6": "516.9",
        "8": "497.2",
        "9": "501.0"
    },
    "8": {
        "0": "492.8",
        "1": "517.0",
        "10": "520.9",
        "11": "510.5",
        "12": "492.7",
        "13": "514.6",
        "14": "501.2",
        "15": "515.4",
        "16": "500.9",
        "17": "523.3",
        "18": "505.3",
        "19": "497.0",
        "2": "495.0",
        "20": "504.2",
        "3": "60.95",
        "4": "520.4",
        "5": "498.5",
        "6": "519.1",
        "7": "505.3",
        "9": "501.4"
    },
    "9": {
        "0": "491.5",
        "1": "516.9",
        "10": "521.0",
        "11": "508.9",
        "12": "494.0",
        "13": "514.6",
        "14": "501.2",
        "15": "515.4",
        "16": "502.1",
        "17": "522.0",
        "18": "506.4",
        "19": "495.6",
        "2": "494.3",
        "20": "504.4",
        "3": "60.93",
        "4": "521.0",
        "5": "498.3",
        "6": "518.6",
        "7": "504.9",
        "8": "497.5"
    }
}

//console.log(endpoint_averages["0"])

function getNeighbors(node) {
  return links.reduce(function (neighbors, link) {
      if (link.target.id === node.id) {
        neighbors.push(link.source.id)
      } else if (link.source.id === node.id) {
        neighbors.push(link.target.id)
      }
      return neighbors
    },
    [node.id]
  )
}

function isNeighborLink(node, link) {
  return link.target.id === node.id || link.source.id === node.id
}


function getNodeColor(node, neighbors) {
  if (Array.isArray(neighbors) && neighbors.indexOf(node.id) > -1) {
    return node.level === 1 ? 'blue' : 'green'
  }
  return node.level === 1 ? 'red' : 'gray'
}

function getDelayToStorageForNode(node, endpoint_averages){
  var delays = [];
  for (var i in endpoint_averages[node]){
    	delays.push(endpoint_averages[node][i]);
  }
  return delays
}

function getLinkColor(node, link) {
  return isNeighborLink(node, link) ? 'green' : '#E5E5E5'
}

function getTextColor(node, neighbors) {
  return Array.isArray(neighbors) && neighbors.indexOf(node.id) > -1 ? 'green' : 'black'
}

var width = 800//window.innerWidth
var height = 600//window.innerHeight

var svg = d3.select('svg')
svg.attr('width', width).attr('height', height)

// simulation setup with all forces
var linkForce = d3
  .forceLink()
  .id(function (link) { return link.id })
  .strength(function (link) { return link.strength })

var simulation = d3
  .forceSimulation()
  .force('link', linkForce)
  .force('charge', d3.forceManyBody().strength(-500))
  .force('center', d3.forceCenter(width / 3.0, height / 2))

var dragDrop = d3.drag().on('start', function (node) {
  node.fx = node.x
  node.fy = node.y
}).on('drag', function (node) {
  simulation.alphaTarget(0.7).restart()
  node.fx = d3.event.x
  node.fy = d3.event.y
}).on('end', function (node) {
  if (!d3.event.active) {
    simulation.alphaTarget(0)
  }
  node.fx = null
  node.fy = null
})

function selectNode(selectedNode) {
  var neighbors = getNeighbors(selectedNode)
  // we modify the styles to highlight selected nodes
  nodeElements.attr('fill', function (node) { return getNodeColor(node, neighbors) })
  textElements.attr('fill', function (node) { return getTextColor(node, neighbors) })
  linkElements.attr('stroke', function (link) { return getLinkColor(selectedNode, link) })
}

var linkElements = svg.append("g")
  .attr("class", "links")
  .selectAll("line")
  .data(links)
  .enter().append("line")
    .attr("stroke-width", 1)
	  .attr("stroke", "rgba(50, 50, 50, 0.2)")

var nodeElements = svg.append("g")
  .attr("class", "nodes")
  .selectAll("circle")
  .data(nodes)
  .enter().append("circle")
    .attr("r", 10)
    .attr("fill", getNodeColor)
    .call(dragDrop)
    .on('click', selectNode)

var textElements = svg.append("g")
  .attr("class", "texts")
  .selectAll("text")
  .data(nodes)
  .enter().append("text")
    .text(function (node) { return  node.label })
	  .attr("font-size", 15)
	  .attr("dx", 15)
    .attr("dy", 4)

// var tableElements = svg.append("g")
// 	.attr("class", "columns")
// 	.selectAll("text")
// 	.data(getDelayToStorageForNode("0", endpoint_averages))
// 	.enter().append("text")
// 	  .attr("font-size", 15)
// 		.attr("fill", "black")
// 	  .attr("dx", 0)
//     .attr("dy", 0)
// 		.text(d => d);

// textPostions = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210]

simulation.nodes(nodes).on('tick', () => {
  nodeElements
    .attr('cx', function (node) { return node.x })
    .attr('cy', function (node) { return node.y })
  textElements
    .attr('x', function (node) { return node.x })
    .attr('y', function (node) { return node.y })
  linkElements
    .attr('x1', function (link) { return link.source.x })
    .attr('y1', function (link) { return link.source.y })
    .attr('x2', function (link) { return link.target.x })
    .attr('y2', function (link) { return link.target.y })
})

// simulation.nodes(nodes).on('tick', () => {
// 	console.log(node.id)
// })
// 	tableElements
//     .attr('x', function (node) { return textPostions })
//     .attr('y', function (node) { return 100 })
// // getNodeDelays('0', endpoint_averages)

simulation.force("link").links(links)
</script>

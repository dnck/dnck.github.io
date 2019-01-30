// Store reference to lists, paragraph and button
const list1 = document.querySelector('.input ul');
const list2 = document.querySelector('.output ul');
const para = document.querySelector('p.stream');
const button = document.querySelector('button');

// create empty string in which to store result
let result = "";

// function to generate random character string

function randomChars() {
let string = "";
let choices = "i merge";

for (let i = 0; i < 6; i++) {
  string += choices.charAt(Math.floor(Math.random() * choices.length));
}
return string;
}

const stream = new ReadableStream({
  start(controller) {
    interval = setInterval(() => {
      let string = randomChars();
      // Add the string to the stream
      controller.enqueue(string);
      // show it on the screen
      let listItem = document.createElement('li');
      listItem.textContent = string;
      list1.appendChild(listItem);}, 250);
    button.addEventListener('click', function(){
      clearInterval(interval);
      readStream();
      controller.close();})
    },
    pull(controller){
      // We don't really need a pull in this example
    },
    cancel(){
      // This is called if the reader cancels,
      // so we should stop generating strings
      clearInterval(interval);
    }
});

function readStream(){
  const reader = stream.getReader();
  let charsReceived = 0;
  // read() returns a promise that resolves when a value has been received
  reader.read().then(function processText({ done, value }) {
    // Result objects contain two properties:
    // done  - true if the stream has already given you all its data.
    // value - some data. Always undefined when done is true.
  if (done){
    console.log("Stream complete");
    para.textContent = result;
    return;
    }
  charsReceived += value.length;
  const chunk = value;
  let listItem = document.createElement('li');
  listItem.textContent = 'Read ' + charsReceived + ' characters so far. Current chunk = ' + chunk;
  list2.appendChild(listItem);
  result += chunk;
  // Read some more, and call this function again
  return reader.read().then(processText);
});
}

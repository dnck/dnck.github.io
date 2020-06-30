<script src="https://d3js.org/d3.v5.min.js"></script>
<script>
  randomSize = function(){
    return Math.floor((Math.random() * 25) + 20);
  };

  randomColor = function(){
    return '#'+(Math.random()*0xFFFFFF<<0).toString(16);
  };

  myfunc = function() {

    var size = 25//randomSize()

    window.onload = function() {
        var dataset;

        dataset = [[0,0], [26,0], [52,0], [78,0], [104,0], [130,0], [156,0], [182,0],
                   [0,26], [26,26], [52,26], [78,26], [104,26], [130,26], [156,26], [182,26],
                   [0,52], [26,52], [52,52], [78,52], [104,52], [130,52], [156,52], [182,52],
                   [0,78], [26,78], [52,78], [78,78], [104,78], [130,78], [156,78], [182,78]];

        var svgContainer = d3.select("#tiles")
                             .append("svg")
                             .attr("viewBox", "0 0 210 100")
                             .attr("xmlns", "http://www.w3.org/2000/svg");

        var rects = svgContainer.selectAll("rect")
                                .data(dataset)
                                .enter()
                                .append("rect");

        var rectsProps = rects.attr("width", size)
                              .attr("height", size)
                              .attr("fill", "#fff")
                              .attr("y", function (d){return d[1];})
                              .attr("x", function (d){return d[0];})

        d3.selectAll("rect").transition()
                            .delay(
                              function(d, i){
                                return i * 100;
                              }
                            )
                            .attr("fill", function(d){
                              return randomColor();
                            })


        // hello = ["H", "E", "L", "L", "O", " ", "W", "O", "R", "L", "D", "!"]
        // var p = d3.select("body")
        //           .selectAll("p")
        //           .data(hello)
        //           .enter()
        //           .append( hello )

    };
  };
myfunc();
</script>

<script src="https://d3js.org/d3.v5.min.js"></script>
<script>
  randomSize = function(){
    return Math.floor((Math.random() * 25) + 20);
  };

  randomColor = function(){
    return '#'+(Math.random()*0xFFFFFF<<0).toString(16);
  };
  degrees_to_radians = function(degrees){
    var pi = Math.PI;
    return degrees * (pi/180);
  }
  radians_to_degrees = function(radians){
    var pi = Math.PI;
    return radians * (180/pi);
  }

  get_x_coords = function(center, angle, offset){
    return center[0] + Math.cos(angle) * offset
  }

  get_y_coords = function(center, angle, offset){
    return center[1] + Math.sin(angle) * offset
  }

  randomPosition = function(){
    x = Math.floor((Math.random() * 25) + 100);
    y = Math.floor((Math.random() * 25) + 100);
    return [get_x_coords([x,y], 45, 20), get_y_coords([x,y], 45, 20)]
  };


  center = [375,125]
  circle0 = [[375,125]]
  var i;
  for (i=270; i>0; i-=45){
    //console.log(get_x_coords(center, i, 50), get_y_coords(center, i, 50))
    circle0.push([get_x_coords(center, i, 90), get_y_coords(center, i, 90)])
    //console.log(get_x_coords(center, i, 100), get_y_coords(center, i, 100))
  };
  for (i=0; i<7; i++){
    for (k=270; k>0; k-=45){
      circle0.push([get_x_coords(circle0[i], k, 30), get_y_coords(circle0[i], k, 30)])
    };
  };

  circle0.reverse()

  myfunc = function(dataset) {

    window.onload = function() {

        var svgContainer = d3.select("#tiles")
                             .append("svg")
                             .attr('width', 800)
                             .attr('height', 300);
                             //.attr("viewBox", "0 0 750 250")
                             //.attr("xmlns", "http://www.w3.org/2000/svg");

        var circles = svgContainer.selectAll("circle")
                                 .data(dataset)
                                 .enter()
                                 .append("circle");

         var circlesProps = circles.attr("r", randomSize())
                           .attr("fill", "transparent")
                           .attr("stroke", "#fff")
                           .attr("cy", function (d){return d[1]})
                           .attr("cx", function (d){return d[0]})

         // d3.selectAll("circle").transition()
         //                     .delay(function(d, i) { return i * 50; })
         //                     .attr("stroke", function(d) {return "#000000";})
         //                     .attr("r", function(d) {return randomSize();})
         //

       d3.selectAll("circle").transition()
                           .delay(function(d, i) { return i * 75; })
                           .attr("fill", function(d) {return randomColor();})
                           .attr("r", function(d) {return "2%";})//randomSize();})


        d3.selectAll("circle").call(d3.drag()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended));

        function dragstarted(d) {
          d3.select(this).raise().classed("active", true);
        }

        function dragged(d) {
          d3.select(this).attr("cx", d.x = d3.event.x).attr("cy", d.y = d3.event.y);
        }

        function dragended(d) {
          d3.select(this).classed("active", false);
        }


        // var buttonContainer = d3.select("#restart")
        //                      .append("button")
        //                      .attr("on click", "/")


    };
  };
myfunc(circle0);
</script>

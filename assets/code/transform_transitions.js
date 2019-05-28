<script src="https://d3js.org/d3.v5.min.js"></script>
<script>
var w = 960,
    h = 500,
    z = 10,
    x = w / z,
    y = h / z;

var svg = d3.select("#tiles")
                     .append("svg")
                     .attr("viewBox", "0 0 750 450")
                     .attr("xmlns", "http://www.w3.org/2000/svg");

svg.selectAll("rect")
    .data(d3.range(x * y))
    .enter().append("rect")
    .attr("transform", translate)
    .attr("width", z)
    .attr("height", z)
    .style("fill", function(d) { return d3.hsl(d % x / x * 360, 1, Math.floor(d / x) / y); })
    .on("mouseover", mouseover);

function translate(d) {
  return "translate(" + (d % x) * z + "," + Math.floor(d / x) * z + ")";
}

function mouseover(d) {
  this.parentNode.appendChild(this);

  d3.select(this)
      .style("pointer-events", "none")
    .transition()
      .duration(750)
      .attr("transform", "translate(480,480)scale(23)rotate(180)")
      
    .transition()
      .delay(250)
      .attr("transform", "translate(240,240)scale(0)")
      .style("fill-opacity", 0)
      .remove();
}
</script>

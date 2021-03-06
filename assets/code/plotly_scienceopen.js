<script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
<script>
(function() {
var d3 = Plotly.d3;
var WIDTH_IN_PERCENT_OF_PARENT = 80,
    HEIGHT_IN_PERCENT_OF_PARENT = 60;
var gd3 = d3.select('h0').append('div').style({
        width: WIDTH_IN_PERCENT_OF_PARENT + '%',
        'margin-left': (100 - WIDTH_IN_PERCENT_OF_PARENT) / 2 + '%',
        height: HEIGHT_IN_PERCENT_OF_PARENT + 'vh',
        'margin-top': (100 - HEIGHT_IN_PERCENT_OF_PARENT) / 2 + 'vh'});
var gd = gd3.node();
		Plotly.d3.csv('https://raw.githubusercontent.com/dnck/dnck.github.io/master/assets/data/test.csv',
			function(err, rows){
				var YEAR = 2016;
				var discipline = ['Discipline A', 'Discipline B', 'Discipline C', 'Discipline D', 'Discipline E'];
				var POP_TO_PX_SIZE = 0.25;
			function unpack(rows, key){
				return rows.map(function(row){
					return row[key];
				});
			}
			var data = discipline.map(function(discipline){
				var rowsFiltered =
				rows.filter(function(row){
					return (row.discipline === discipline) && (+row.year === YEAR);
				});
						return{
								mode: 'markers',
								name: discipline,
								x: unpack(rowsFiltered, 'readCount'),
								y: unpack(rowsFiltered, 'citedBy'),
								text: unpack(rowsFiltered, 'contentTitle'),
								marker:{sizemode: 'area', size: unpack(rowsFiltered, 'readCount'), sizeref: POP_TO_PX_SIZE},
						};
			});
			var layout = {
				title: 'Cited-by Count ~ Read Count',
				xaxis: {title: 'Read Count'},
				yaxis: {title: 'Cited by Count Total'},
				margin: {t: 50},
				autosize: true,
				hovermode: 'closest',
			};
			Plotly.plot(gd, data, layout, {showLink: false});
		});
window.onresize = function() {
    Plotly.Plots.resize(gd);
};
})();
</script>

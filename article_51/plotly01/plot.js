var graphElement = document.getElementById('graphElement');
var data = {
    x: [1, 2, 3, 4, 5],
    y: [1, 2, 4, 8, 16]
}
var opts = {
    margin: { t: 0}
}
Plotly.plot(graphElement, [data], opts)

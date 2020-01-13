$(document).ready(function() {
    $.ajaxSetup({
        async: false
    });

    var series1, series2, series3;

    $.getJSON("/series?offset=0",
        function (data) {
            series1 = data;
            series1.mode = 'lines'
        });
    $.getJSON("/series?offset=-3.14",
        function (data) {
            series2 = data;
            series2.mode = 'markers'
        });
    $.getJSON("/series?offset=0.25",
        function (data) {
            series3 = data;
            series3.mode = 'lines+markers'
        });
    var graphElement = $('#graphElement')[0];
    var opts = {
        margin: { t: 0}
    };
    Plotly.plot(graphElement, [series1, series2, series3], opts);
});

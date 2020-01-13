$(document).ready(function() {
    $.ajaxSetup({
        async: false
    });

    var series1, series2;

    $.getJSON("/series?offset=-3.14",
        function (data) {
            series1 = data;
            series1.type = 'scatter';
        });
    $.getJSON("/series?offset=0",
        function (data) {
            series2 = data;
            series2.type = 'bar';
        });
    var graphElement = $('#graphElement')[0];
    var opts = {
        margin: { t: 0}
    };
    Plotly.plot(graphElement, [series1, series2], opts);
});

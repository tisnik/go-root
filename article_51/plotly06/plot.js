$(document).ready(function() {
    $.ajaxSetup({
        async: false
    });

    var series1, series2, series3;

    $.getJSON("/series1",
        function (data) {
            series1 = data;
            console.log(series1);
        });
    $.getJSON("/series2",
        function (data) {
            series2 = data;
            console.log(series2);
        });
    $.getJSON("/series3",
        function (data) {
            series3 = data;
            console.log(series3);
        });
    var graphElement = $('#graphElement')[0];
    var opts = {
        margin: { t: 0}
    };
    Plotly.plot(graphElement, [series1, series2, series3], opts);
});

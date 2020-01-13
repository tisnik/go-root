$(document).ready(function() {
    $.getJSON("/data",
        function (data) {
            var graphElement = $('#graphElement')[0];
            var opts = {
                margin: { t: 0}
            };
            Plotly.plot(graphElement, [data], opts);
        });
});

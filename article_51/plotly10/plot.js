$(document).ready(function() {
    $.getJSON("/values",
        function (data) {
            var graphElement = $('#graphElement')[0];
            var pie = {};
            pie.values = data;
            pie.labels = ['Go', 'Rust', 'C', 'Java'];
            pie.type = 'pie';
            pie.hole = 0.4;
            var opts = {
                margin: { t: 0}
            };
            Plotly.plot(graphElement, [pie], opts);
        });
});

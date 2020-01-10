$(document).ready(function() {
    $.getJSON("data.json",
        function (data) {
            var graphElement = document.getElementById('graphElement');
            var opts = {
                margin: { t: 0}
            };
            Plotly.plot(graphElement, [data], opts);
        });
});

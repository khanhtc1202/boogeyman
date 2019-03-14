function query() {
	$("#search-form").submit(function( event ) {
		var query = $("#search-form input[name=q]").val();
		var engine = $("#search-form select[name=e]").val();
		var strategy = $("#search-form select[name=s]").val();

		var url = "http://localhost:3000/search?q="+query+"&s="+strategy+"&e="+engine;
		$.get(url, function (data, status) {
			console.log(status);
			console.log(data);
		});

		event.preventDefault();
	});
}

// function renderResult(data) {
// 	var json = $.parseJSON(data);
// 	$.each(json["results"], function (result) {
//
// 	})
// }

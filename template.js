// function fill_template(){
	
	var data={
		title: "test",
		test : "<span style='height: 14px;'>Название</span> <input style='' type='text' value='' />",
		NewProfileList1:[
			"test"
			
			
		],
		NewProfileList:[
		{"type": 1}
		
		],
		
		footer: "footer"
	};
	var template = Handlebars.compile(document.querySelector("#template").innerHTML);
	var filled = template(data, {
			noEscape: true,
		}
		
	);
	document.querySelector("#output").innerHTML = filled;
	
	
// }
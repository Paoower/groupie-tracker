const data = {
    fromCreationYear: 1950,
    toCreationYear: 2024,
    fromAlbumYear: 1950,
    toAlbumYear: 2024,
    membersMin: 1,
    membersMax: 8,
    location: ""
}

data.fromCreationYear = 1950
data.toCreationYear = 2024
data.fromAlbumYear = 1950
data.toAlbumYear = 2024

const handleFilter = async () => {
	reqOptions = {
		method: "POST",
		body: JSON.stringify(data)
	};
	const req = await fetch("/filters", reqOptions);
	const text = await req.text();
	let container = document.getElementById("artists-container").innerHTML
	if (text != container) {
		document.getElementById("artists-container").innerHTML = text;
	}
}

document.querySelectorAll("input[type=range]").forEach(range => {
	range.addEventListener("input", async e => {
		if (e.target.id == "slider-1") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[1];
			if (parseInt(e.target.value) > parseInt(s.value)) {
				e.target.value = s.value;
			}
		}
		if (e.target.id == "slider-2") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) < parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		if (e.target.id == "slider-3") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[1];
			if (parseInt(e.target.value) > parseInt(s.value)) {
				e.target.value = s.value;
			}
		}
		if (e.target.id == "slider-4") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) < parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		if (e.target.id == "slider-5") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) > parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		if (e.target.id == "slider-6") {
			let s = e.target.closest("div").querySelectorAll("input[type=range]")[0];
			if (parseInt(e.target.value) < parseInt(s.value)) {
				e.target.value = s.value;
			}
		}

		data[e.target.name] = parseInt(e.target.value);
		labelsDiv = e.target.closest("div").querySelector("div");
		labelsDiv.querySelectorAll("label").forEach(label => {
			if (label.htmlFor == e.target.id) {
				label.innerHTML = e.target.value
			}
		})
		await handleFilter();
	})
})

document.querySelectorAll("select").forEach(select => {
	select.addEventListener("change", e => {
		if (e.target.name == "members") {
			data[e.target.name] = parseInt(e.target.value)
			handleFilter()
		} else if (e.target.name == "location") {
			data[e.target.name] = e.target.value
			handleFilter()
		} else {
			data[e.target.name] = e.target.value
			handleFilter()
		}
	})
})
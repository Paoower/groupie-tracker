const linkClasses = [
    "p-1.5", "pl-2", "cursor-pointer", "transition-all",
    "text-sm", "hover:bg-gray-100", "w-full"
];

const searchFetch = async (value) => {
    const reqOptions = {
        method: "POST",
        body: JSON.stringify({ search: value })
    };
    const req = await fetch("/searchbar", reqOptions);
    return await req.json();
};

const searchLink = (content, href) => {
    const li = document.createElement("li");
    linkClasses.forEach(c => li.classList.add(c));
    const a = document.createElement("a");
    a.innerHTML = content;
    a.href = href;
    a.classList.add("block");
    li.appendChild(a);
    return li;
};

export function initializeSearch() {
    const searchInput = document.getElementById("search");
    const resultsContainer = document.getElementById("results");

    if (searchInput && resultsContainer) {
        let debounceTimer;

        searchInput.addEventListener("input", async (e) => {
            resultsContainer.textContent = "";
            if (e.target.value.trim().length !== 0) {
                const json = await searchFetch(e.target.value);
                if (json.artists || json.members || json.locations || json.firstAlbumDates || json.creationDates) {
                    resultsContainer.classList.remove("hidden");

                    // Display artist results
                    json.artists?.forEach(artist => {
                        const content = `${artist.name} - artist/band`;
                        const href = `http://localhost:8080/artist?id=${artist.id}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });

                    // Display member results
                    json.members?.forEach(member => {
                        const content = `${member.name} - ${member.type}${member.artist.name}`;
                        const href = `http://localhost:8080/artist?id=${member.artist.id}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });

                    // Display location results
                    json.locations?.forEach(location => {
                        const content = `${location.artist} - concert location in ${location.name}`;
                        const href = `http://localhost:8080/artist?id=${location.artistId}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });

                    // Display first album date results
                    json.firstAlbumDates?.forEach(album => {
                        const content = `${album.date} - first album date of ${album.artist}`;
                        const href = `http://localhost:8080/artist?id=${album.artistId}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });

                    // Display creation date results
                    json.creationDates?.forEach(creation => {
                        const content = `${creation.date} - creation date of ${creation.artist}`;
                        const href = `http://localhost:8080/artist?id=${creation.artistId}`;
                        resultsContainer.appendChild(searchLink(content, href));
                    });
                } else {
                    resultsContainer.classList.add("hidden");
                }
            } else {
                resultsContainer.classList.add("hidden");
            }
        });

        // Close results when clicking outside
        document.addEventListener('click', (e) => {
            if (!searchInput.contains(e.target) && !resultsContainer.contains(e.target)) {
                resultsContainer.classList.add("hidden");
            }
        });

        // Keyboard navigation
        searchInput.addEventListener('keydown', (e) => {
            const items = resultsContainer.querySelectorAll('li');
            let currentIndex = Array.from(items).findIndex(item => item.classList.contains('bg-gray-200'));

            switch (e.key) {
                case 'ArrowDown':
                    e.preventDefault();
                    if (currentIndex < items.length - 1) currentIndex++;
                    break;
                case 'ArrowUp':
                    e.preventDefault();
                    if (currentIndex > 0) currentIndex--;
                    break;
                case 'Enter':
                    e.preventDefault();
                    if (currentIndex !== -1) {
                        items[currentIndex].querySelector('a').click();
                    }
                    return;
                default:
                    return;
            }

            items.forEach(item => item.classList.remove('bg-gray-200'));
            if (items[currentIndex]) {
                items[currentIndex].classList.add('bg-gray-200');
                items[currentIndex].scrollIntoView({ block: 'nearest' });
            }
        });
    }
}


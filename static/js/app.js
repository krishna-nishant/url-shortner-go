document.addEventListener('DOMContentLoaded', function() {
    const urlForm = document.getElementById('url-form');
    const originalUrlInput = document.getElementById('original-url');
    const resultDiv = document.getElementById('result');
    const shortenedUrlInput = document.getElementById('shortened-url');
    const copyBtn = document.getElementById('copy-btn');
    const urlsTableBody = document.getElementById('urls-table-body');

    // Load existing URLs
    loadUrls();

    // Handle form submission
    urlForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        const originalUrl = originalUrlInput.value.trim();
        if (!originalUrl) return;

        // Check if URL has http/https prefix
        let formattedUrl = originalUrl;
        if (!originalUrl.startsWith('http://') && !originalUrl.startsWith('https://')) {
            formattedUrl = 'https://' + originalUrl;
        }

        // Send API request
        fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ original: formattedUrl })
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert('Error: ' + data.error);
                return;
            }

            // Display the shortened URL
            const shortUrl = window.location.origin + '/' + data.short;
            shortenedUrlInput.value = shortUrl;
            resultDiv.classList.remove('d-none');
            
            // Reload URLs list
            loadUrls();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while shortening the URL.');
        });
    });

    // Handle copy button
    copyBtn.addEventListener('click', function() {
        shortenedUrlInput.select();
        document.execCommand('copy');
        
        // Show copy success feedback
        const originalText = copyBtn.textContent;
        copyBtn.textContent = 'Copied!';
        copyBtn.classList.add('copy-success');
        
        setTimeout(() => {
            copyBtn.textContent = originalText;
            copyBtn.classList.remove('copy-success');
        }, 2000);
    });

    // Function to load URLs from API
    function loadUrls() {
        fetch('/api/urls')
            .then(response => response.json())
            .then(data => {
                // Clear table
                urlsTableBody.innerHTML = '';
                
                // Add URLs to table
                data.forEach(url => {
                    const row = document.createElement('tr');
                    
                    // Original URL column
                    const originalCell = document.createElement('td');
                    const originalLink = document.createElement('a');
                    originalLink.href = url.original;
                    originalLink.target = '_blank';
                    originalLink.classList.add('url-text');
                    originalLink.textContent = url.original;
                    originalLink.title = url.original;
                    originalCell.appendChild(originalLink);
                    
                    // Short URL column
                    const shortCell = document.createElement('td');
                    const shortUrl = window.location.origin + '/' + url.short;
                    const shortLink = document.createElement('a');
                    shortLink.href = shortUrl;
                    shortLink.target = '_blank';
                    shortLink.textContent = url.short;
                    shortCell.appendChild(shortLink);
                    
                    // Click count column
                    const clicksCell = document.createElement('td');
                    clicksCell.textContent = url.click_count;
                    
                    // Add cells to row
                    row.appendChild(originalCell);
                    row.appendChild(shortCell);
                    row.appendChild(clicksCell);
                    
                    // Add row to table
                    urlsTableBody.appendChild(row);
                });
            })
            .catch(error => {
                console.error('Error loading URLs:', error);
            });
    }
}); 
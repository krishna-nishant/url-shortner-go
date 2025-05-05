document.addEventListener("DOMContentLoaded", () => {
    const urlForm = document.getElementById("url-form")
    const originalUrlInput = document.getElementById("original-url")
    const resultDiv = document.getElementById("result")
    const shortenedUrlInput = document.getElementById("shortened-url")
    const copyBtn = document.getElementById("copy-btn")
    const urlsTableBody = document.getElementById("urls-table-body")
  
    // Load existing URLs
    loadUrls()
  
    // Handle form submission
    urlForm.addEventListener("submit", (e) => {
      e.preventDefault()
  
      const originalUrl = originalUrlInput.value.trim()
      if (!originalUrl) return
  
      // Check if URL has http/https prefix
      let formattedUrl = originalUrl
      if (!originalUrl.startsWith("http://") && !originalUrl.startsWith("https://")) {
        formattedUrl = "https://" + originalUrl
      }
  
      // Send API request
      fetch("/shorten", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ original: formattedUrl }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.error) {
            alert("Error: " + data.error)
            return
          }
  
          // Display the shortened URL
          const shortUrl = window.location.origin + "/" + data.short
          shortenedUrlInput.value = shortUrl
          resultDiv.classList.remove("hidden")
  
          // Reload URLs list
          loadUrls()
        })
        .catch((error) => {
          console.error("Error:", error)
          alert("An error occurred while shortening the URL.")
        })
    })
  
    // Handle copy button
    copyBtn.addEventListener("click", () => {
      shortenedUrlInput.select()
      document.execCommand("copy")
  
      // Show copy success feedback
      const originalHTML = copyBtn.innerHTML
      copyBtn.innerHTML = `
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
          `
      copyBtn.classList.add("copy-success")
  
      setTimeout(() => {
        copyBtn.innerHTML = originalHTML
        copyBtn.classList.remove("copy-success")
      }, 2000)
    })
  
    // Function to load URLs from API
    function loadUrls() {
      fetch("/api/urls")
        .then((response) => response.json())
        .then((data) => {
          // Clear table
          urlsTableBody.innerHTML = ""
  
          // Add URLs to table
          data.forEach((url) => {
            const row = document.createElement("tr")
            row.className = "hover:bg-slate-50"
  
            // Original URL column
            const originalCell = document.createElement("td")
            originalCell.className = "py-3 px-4"
            const originalLink = document.createElement("a")
            originalLink.href = url.original
            originalLink.target = "_blank"
            originalLink.className = "url-text text-violet-600 hover:text-violet-800 transition-colors"
            originalLink.textContent = url.original
            originalLink.title = url.original
            originalCell.appendChild(originalLink)
  
            // Short URL column
            const shortCell = document.createElement("td")
            shortCell.className = "py-3 px-4"
            const shortUrl = window.location.origin + "/" + url.short
            const shortLink = document.createElement("a")
            shortLink.href = shortUrl
            shortLink.target = "_blank"
            shortLink.className = "font-medium text-slate-800 hover:text-violet-600 transition-colors"
            shortLink.textContent = url.short
            shortCell.appendChild(shortLink)
  
            // Click count column
            const clicksCell = document.createElement("td")
            clicksCell.className = "py-3 px-4"
            const clicksBadge = document.createElement("span")
            clicksBadge.className =
              "inline-flex items-center justify-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-slate-100 text-slate-800"
            clicksBadge.textContent = url.click_count
            clicksCell.appendChild(clicksBadge)
  
            // Add cells to row
            row.appendChild(originalCell)
            row.appendChild(shortCell)
            row.appendChild(clicksCell)
  
            // Add row to table
            urlsTableBody.appendChild(row)
          })
        })
        .catch((error) => {
          console.error("Error loading URLs:", error)
        })
    }
  })
  
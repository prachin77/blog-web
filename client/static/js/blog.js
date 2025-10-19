// Handle comment form submission
document.querySelector('.comment-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const textarea = this.querySelector('.comment-textarea');
    if (textarea.value.trim()) {
        alert('Comment posted! (This is a demo)');
        textarea.value = '';
    }
});

// Handle share buttons
document.querySelectorAll('.share-btn').forEach(btn => {
    btn.addEventListener('click', function() {
        const platform = this.classList.contains('share-twitter') ? 'Twitter' :
                       this.classList.contains('share-linkedin') ? 'LinkedIn' :
                       this.classList.contains('share-facebook') ? 'Facebook' : 'Clipboard';
        
        if (platform === 'Clipboard') {
            navigator.clipboard.writeText(window.location.href);
            alert('Link copied to clipboard!');
        } else {
            alert(`Share on ${platform} (This is a demo)`);
        }
    });
});

// Handle tag clicks
document.querySelectorAll('.tag').forEach(tag => {
    tag.addEventListener('click', function() {
        alert(`Filter by ${this.textContent} (This is a demo)`);
    });
});

// Initialize related post images with placeholders
document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.related-image').forEach(img => {
        if (!img.src || img.src === window.location.href) {
            img.style.display = 'none';
            const placeholder = document.createElement('div');
            placeholder.className = 'related-image-placeholder';
            placeholder.textContent = '📝';
            img.parentElement.appendChild(placeholder);
        }
    });
});
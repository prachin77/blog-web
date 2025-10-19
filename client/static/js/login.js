// Form submission with animation
document.getElementById('signinForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const button = this.querySelector('.btn-signin');
    const originalText = button.textContent;
    
    // Show loading state
    button.textContent = 'Signing In...';
    button.style.pointerEvents = 'none';
    button.style.opacity = '0.8';
    
    // Simulate sign in
    setTimeout(() => {
        button.classList.add('success');
        button.textContent = '✓ Welcome Back!';
        
        // Redirect after success animation
        setTimeout(() => {
            window.location.href = 'index.html';
        }, 1000);
    }, 1500);
});

// Input validation feedback
const inputs = document.querySelectorAll('.input-field');
inputs.forEach(input => {
    input.addEventListener('focus', function() {
        this.style.transform = 'scale(1.01)';
    });
    
    input.addEventListener('blur', function() {
        this.style.transform = 'scale(1)';
        
        // Simple validation feedback
        if (this.value && this.checkValidity()) {
            this.style.borderColor = '#28a745';
        } else if (this.value) {
            this.style.borderColor = '#dc3545';
        }
    });

    input.addEventListener('input', function() {
        if (this.style.borderColor === 'rgb(220, 53, 69)' || this.style.borderColor === 'rgb(40, 167, 69)') {
            this.style.borderColor = '#e0e0e0';
        }
    });
});

// Social button animations
const socialButtons = document.querySelectorAll('.btn-social');
socialButtons.forEach(button => {
    button.addEventListener('click', function(e) {
        e.preventDefault();
        this.style.transform = 'scale(0.95)';
        setTimeout(() => {
            this.style.transform = 'scale(1)';
        }, 100);
    });
});

// Forgot password link
document.querySelector('.forgot-password').addEventListener('click', function(e) {
    e.preventDefault();
    alert('Password reset functionality would be implemented here. An email would be sent to reset your password.');
});
// Form submission with animation
document.getElementById('signupForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const button = this.querySelector('.btn-signup');
    const originalText = button.textContent;
    
    // Show loading state
    button.textContent = 'Creating Account...';
    button.style.pointerEvents = 'none';
    button.style.opacity = '0.8';
    
    // Simulate account creation
    setTimeout(() => {
        button.classList.add('success');
        button.textContent = '✓ Account Created!';
        
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

// Password strength indicator
const passwordInput = document.getElementById('password');
passwordInput.addEventListener('input', function() {
    const strength = calculatePasswordStrength(this.value);
    // You can add visual feedback here if desired
});

function calculatePasswordStrength(password) {
    let strength = 0;
    if (password.length >= 8) strength++;
    if (password.match(/[a-z]/) && password.match(/[A-Z]/)) strength++;
    if (password.match(/[0-9]/)) strength++;
    if (password.match(/[^a-zA-Z0-9]/)) strength++;
    return strength;
}

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
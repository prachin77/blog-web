// Navbar scroll effect
window.addEventListener('scroll', () => {
    const navbar = document.querySelector('.navbar');
    if (window.scrollY > 50) {
        navbar.classList.add('scrolled');
    } else {
        navbar.classList.remove('scrolled');
    }
});

// Parallax effect for hero section
let heroSection = document.querySelector('.hero');
let heroContent = document.querySelector('.hero-content');
let heroAfter = document.querySelector('.hero::after');

window.addEventListener('scroll', () => {
    let scrolled = window.pageYOffset;
    let heroHeight = heroSection.offsetHeight;
    
    if (scrolled < heroHeight) {
        // Parallax for hero content
        heroContent.style.transform = `translateY(${scrolled * 0.5}px)`;
        heroContent.style.opacity = 1 - (scrolled / heroHeight) * 0.8;
        
        // Parallax for hero background shapes
        let heroBeforeEl = document.querySelector('.hero::before');
        heroSection.style.backgroundPositionY = `${scrolled * 0.3}px`;
    }
});

// Parallax effect for blog cards on scroll
const blogCards = document.querySelectorAll('.blog-card');
const sectionBefore = document.querySelector('.section::before');

window.addEventListener('scroll', () => {
    let scrolled = window.pageYOffset;
    
    // Parallax for section decorative element
    if (sectionBefore) {
        let sectionOffset = document.querySelector('.section').offsetTop;
        let relativeScroll = scrolled - sectionOffset;
        if (relativeScroll > -500 && relativeScroll < 1000) {
            document.querySelector('.section').style.setProperty('--parallax-offset', `${relativeScroll * 0.1}px`);
        }
    }

    // Stagger animation for blog cards
    blogCards.forEach((card, index) => {
        let cardTop = card.getBoundingClientRect().top;
        let cardBottom = card.getBoundingClientRect().bottom;
        let windowHeight = window.innerHeight;
        
        if (cardTop < windowHeight && cardBottom > 0) {
            let scrollPercent = (windowHeight - cardTop) / windowHeight;
            let translateY = (1 - scrollPercent) * 30;
            
            if (translateY > 0) {
                card.style.transform = `translateY(${translateY}px)`;
                card.style.opacity = scrollPercent;
            } else {
                card.style.transform = 'translateY(0)';
                card.style.opacity = 1;
            }
        }
    });
});

// Enhanced parallax for hero background elements
const heroBeforeElement = window.getComputedStyle(heroSection, '::before');
const heroAfterElement = window.getComputedStyle(heroSection, '::after');

window.addEventListener('scroll', () => {
    let scrolled = window.pageYOffset;
    
    // Create dynamic parallax effect using CSS variables
    document.documentElement.style.setProperty('--scroll-y', `${scrolled * 0.5}px`);
    document.documentElement.style.setProperty('--scroll-y-slow', `${scrolled * 0.2}px`);
});

// Add CSS variable support for parallax
const style = document.createElement('style');
style.textContent = `
    .hero::before {
        transform: translateY(var(--scroll-y-slow, 0)) translate(-50px, 50px) scale(1.1) !important;
    }
    .hero::after {
        transform: translateY(var(--scroll-y, 0)) !important;
    }
    .section::before {
        transform: translateY(var(--parallax-offset, 0)) !important;
    }
`;
document.head.appendChild(style);

// Smooth scroll
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({ behavior: 'smooth' });
        }
    });
});

// Initial animation on page load
window.addEventListener('load', () => {
    heroContent.style.transition = 'transform 0.8s ease-out, opacity 0.8s ease-out';
    heroContent.style.transform = 'translateY(0)';
    heroContent.style.opacity = '1';

    // Animate blog cards in sequence
    blogCards.forEach((card, index) => {
        setTimeout(() => {
            card.style.transition = 'all 0.6s ease-out';
            card.style.transform = 'translateY(0)';
            card.style.opacity = '1';
        }, index * 100);
    });
});

// Set initial state for cards
blogCards.forEach(card => {
    card.style.opacity = '0';
    card.style.transform = 'translateY(30px)';
});
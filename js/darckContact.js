function googleTranslateElementInit(){
	new google.translate.TranslateElement({pageLanguaje: 'es', layout: google.translate.TranslateElement.InlineLayout.SIMPLE}, 'google_translate_element');
};

const btnSwitch = document.querySelector('#switch');
var card = document.querySelector('.card');



btnSwitch.addEventListener('click', () => {
	document.body.classList.toggle('dark');
	card.classList.toggle('cardDark');
	btnSwitch.classList.toggle('active');

	// Guardamos el modo en localstorage.
	if(document.body.classList.contains('dark')){
		localStorage.setItem('dark-mode', 'true');
	} else {
		localStorage.setItem('dark-mode', 'false');
	}
});

// Obtenemos el modo actual.
if(localStorage.getItem('dark-mode') === 'true'){
	document.body.classList.add('dark');
	card.classList.add('cardDark');
	btnSwitch.classList.add('active');
} else {
	document.body.classList.remove('dark');
	card.classList.remove('cardDark');
	btnSwitch.classList.remove('active');
};



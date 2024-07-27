/**
  * @param {Element} element 
  * @param {() => any} element 
  */
function addOnClickOutsideListener(element, callback) {
  document.addEventListener('click', event => {
    if (!element.contains(event.target)) {
      callback()
    }
  })
}

/**
  * @param {string} id 
  */
function hide(id) {
  document.getElementById(id).style.display = "none"
}

/**
  * @param {string} id 
  */
function show(id) {
  document.getElementById(id).style.display = "block"
}

/**
  * @param {string} id 
  */
function toggle(id) {
  if (isOpen(id)) {
    hide(id)
  } else {
    show(id)
  }
}

/**
  * @param {string} id 
  */
function isOpen(id) {
  const elm = document.getElementById(id)
  return !elm.hidden && elm.style.display !== "none"
}

function handleError(event) {
  const response = event.detail.xhr.response
  alert(response)
}

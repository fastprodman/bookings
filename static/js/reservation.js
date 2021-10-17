let noty = document.getElementById('notification').addEventListener("click", function () {
    let html = `
    <form action="reservation.html" method="get" data-date="Today" id="range-picker1" class="needs-validation" novalidate>
                    <div class="input-group has-validation">
                        <span class="input-group-text">From</span>
                        <input type="text" autocomplete="off" id="start" placeholder="Enter starting date" name="start-date" aria-label="First day"
                            class="form-control" required>
                        <span class="input-group-text">to</span>
                        <input type="text"  autocomplete="off" id="end" placeholder="Enter ending date" name="end-date" aria-label="Last day"
                            class="form-control" required>
                        <div class="valid-feedback">
                            Looks good!
                        </div>
                        <div class="invalid-feedback">
                            Please choose a day.
                        </div>
                    </div>
                    <div class="row m-2">
                        <button type="submit" class="btn btn-outline-success btn-block mt-2">Search</button>
                    </div>
                </form>
                
    `
    attention.custome({content: html, title: 'Pick a date'});
})

// Module code

let attention = prompt()

function prompt() {
    let toast = function (c) {
        const {
            icon = "success",
            title = "",
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            icon: icon,
            title: title,
            position: 'top-end',
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function(c){
        const {
            title = "",
            text = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'success',
            title: title,
            text: text,
            footer: footer
        })
    }

    let error = function(c){
        const {
            title = "",
            text = "",
            footer = "",
        } = c;

        Swal.fire({
            icon: 'error',
            title: title,
            text: text,
            footer: footer
        })
    }

        async function custome(c){
        const {
            content = "",
            title = "",
        } = c;

        const { value: formValues } = await Swal.fire({
            title: title,
            html: content,
            focusConfirm: false,
            showCloseButton: true,
            showCancleButton: true,
            showConfirmButton: true,
            willOpen: () => {
                const elem1 = document.getElementById('range-picker1');
                const rp = new DateRangePicker(elem1, {
                format: "dd-mm-yyyy",
                showOnFocus: true,
            });
            },

            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })
        
        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custome: custome,
    }
}
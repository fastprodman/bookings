
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

        const { value: result } = await Swal.fire({
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
        
        if(result){
            if(result.dismiss !== Swal.DismissReason.cancel){
                if(result.value !== ""){
                    if(c.callback !== undefined){
                        c.callback(result);
                    }
                }else{
                    c.callback(false);
                }
            }else{
                c.callback(false);
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custome: custome,
    }
}
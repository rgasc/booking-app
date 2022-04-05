document
    .getElementById("check-availability")
    .addEventListener("click", function () {
        const modalForm = `
                <form id="check-availability-form" action="" method="POST" class="needs-validation" novalidate>
                    <div class="row">
                        <div class="col">
                            <div id="reservation-dates-modal" class="row">
                                <div class="col">
                                    <input disabled required type="text" class="form-control" name="start" id="start" placeholder="Arrival">
                                </div>
                                <div class="col">
                                    <input disabled required type="text" class="form-control" name="end" id="end" placeholder="Departure">
                                </div>
                            </div>
                        </div>
                    </div>
                </form>
            `;

        attention.custom({
            msg: modalForm,
            title: "Choose your dates",
            willOpen: () => {
                const elem = document.getElementById("reservation-dates-modal");
                const rp = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                    showOnFocus: true,
                });
            },
            didOpen: () => {
                document.getElementById("start").removeAttribute("disabled");
                document.getElementById("end").removeAttribute("disabled");
            },
            callback: (result) => {
                const form = document.getElementById("check-availability-form");
                const formData = new FormData(form);
                formData.append("csrf_token", token);

                fetch("/search-availability-json", {
                    method: "POST",
                    body: formData,
                })
                    .then((response) => response.json())
                    .then((data) => {
                        console.log(data.ok);
                        console.log(data.message);
                    });
            },
        });
    });

var tips = {
    getFormValus: function (formId) {
        var data = {};
        var formVal = $(formId).serializeArray();
        var temp = {};
        $.each(formVal, function (inx, val) {
            temp[val.name] = val.value;
            $.extend(data, temp);
        });
        return data;
    },
    ajax: function () {
        $.ajax({
            type: '',
            url: '',
            data: {},
            success: function () {

            },
            error: function () {

            }
        });
    }
};

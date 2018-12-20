var tips = {
    getFormValues: function (formId) {
        var data = {};
        var formVal = $(formId).serializeArray();
        var temp = {};
        $.each(formVal, function (inx, val) {
            temp[val.name] = val.value;
            $.extend(data, temp);
        });
        return data;
    },
    ajax: function (config) {


        if (!config.hasOwnProperty('type')) {
            config.type = 'get';
        }


        if (!config.hasOwnProperty('url')) {
            config.url = '';
        }

        if (!config.hasOwnProperty('data')) {
            config.data = {};
        }

        if (!config.hasOwnProperty('success')) {
            config.success = function (result) {
                layer.msg(result.message, {icon: 6, time: 3000});
            };
        }


        $.ajax({
            type: config.type,
            url: config.url,
            data: config.data,
            success: config.success,
            error: function () {
                layer.alert("异步请求失败", {icon: 5});
            }
        });
    },
    bootstrapTable: function (config) {
        if (typeof config !== 'object') {
            layer.alert("参数必须为js对象", {icon: 5});
            return;
        }

        if (!config.hasOwnProperty('url')) {
            config.url = '';
        }
        if (!config.hasOwnProperty('picker')) {
            config.picker = '#table';
        }

        if (!config.hasOwnProperty('columns')) {
            config.columns = [];
        }

        if (!config.hasOwnProperty('queryParams')) {
            config.queryParams = function (params) {
                params.page = params.pageNumber;
                params.size = params.pageSize;
                $.extend(params, {keyWord: $("#keyWord").val()});
                return params;
            }
        }

        if (!config.hasOwnProperty('pageSize')) {
            config.pageSize = 10;
        }
        $(config.picker).bootstrapTable({
            classes: 'table table-hover', //bootstrap的表格样式
            sidePagination: 'server', //获取数据方式【从服务器获取数据】
            url: config.url,//ajax链接
            pagination: true, //分页
            pageNumber: 1, //页码【第X页】
            pageSize: 10, //每页显示多少条数据
            sortName: 'id', //排序字段
            sortOrder: 'DESC',//排序方式
            queryParamsType: '', // limit
            striped: true,
            queryParams: config.queryParams,
            columns: config.columns

        });
    },
    tableRefresh: function (picker) {
        if (picker === undefined) {
            picker = '#table';
        }
        $(picker).bootstrapTable('refresh');
    }
};

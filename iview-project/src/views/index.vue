<template>
    <div>
        <Menu mode="horizontal" theme="dark" active-name="1">
            <MenuItem name="1" to="/">
                <Icon type="ios-paper"/>
                Asciinema-viewer
            </MenuItem>
        </Menu>
        <br>
        <Upload
                multiple
                type="drag"
                name="myfile"
                :show-upload-list=false
                :on-success="handleRead"
                action="/file/Upload">
            <div style="padding: 20px 0">
                <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                <p>Click or drag files here to upload</p>
            </div>
        </Upload>
        <Table :columns="columns1" :data="data1"></Table>
    </div>
</template>
<script>
    import Util from '../libs/util';

    export default {
        data() {
            return {
                columns1: [
                    {
                        title: 'List',
                        key: 'Path',
                        sortable: true,
                        sortType: "asc",
                        render: (h, params) => {
                            return h('div', [
                                h('a', {
                                    on: {
                                        click: () => {
                                            window.open("viewer?path=" + params.row.Path);
                                        }
                                    }
                                }, params.row.Path)
                            ]);
                        }
                    }
                ],
                data1: []
            }
        }, mounted() {
            this.handleRead();
        }, methods: {
            handleRead() {
                Util.ajax.get("list").then((response) => {
                        console.log(response);
                        this.data1 = response.data;
                    }
                ).catch(function (error) {
                    console.log(error);
                });
            }, Trim(str) {
                return str.replace(/(^\s*)|(\s*$)/g, "");
            }
        }
    }
</script>

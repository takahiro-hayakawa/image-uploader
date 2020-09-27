<template>
  <div class="drop_area" @dragover.prevent @drop.prevent="dropFile">
    ファイルアップロード
  </div>
</template>

<script>
import * as axios from "axios";
import STATUS from '../constants/upload_status';

export default {
  name: 'Upload',
  data() {
    return {
      files: []
    }
  },
  methods: {
    dropFile(event) {
      this.$emit("emitUp", STATUS.UPLOADING)

      this.files = [...event.dataTransfer.files]
      this.files.forEach(file => {
        let form = new FormData()
        form.append('file', file)
        axios({
          method: 'post',
          url :'http://localhost:8000/upload',
          data : form,
          header :{
            'Content-Type': 'multipart/form-data',
          }
        }).then(response => {
          console.log(response.data)
          this.$emit("emitUp", STATUS.UPLOADED)
        }).catch(error => {
          console.log(error)
          this.$emit("emitUp", STATUS.UPLOAD)
        })
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.drop_area {
  position: absolute;
  width: 338px;
  height: 218.9px;
  left: 551.82px;
  top: 427.97px;
  background: #F6F8FB;
  border: 1px dashed #97BEF4;
  box-sizing: border-box;
  border-radius: 12px;
}

.enter {
  border: 10px dotted powderblue;
}
</style>

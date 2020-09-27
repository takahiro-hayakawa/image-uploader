<template>
  <div id="app">
    <h1 class="title">Upload your image</h1>
    <h2 class="description">File should be Jpeg,Png,...</h2>
    <Upload v-show="upload" @emitUp="changeStatus"/>
    <Uploading v-show="uploading"/>
    <Uploaded v-show="uploaded"/>
  </div>
</template>

<script>
import Upload from './components/Upload.vue'
import Uploading from './components/Uploading.vue'
import Uploaded from './components/Uploaded.vue';
import STATUS from './constants/upload_status';

export default {
  name: 'App',
  components: {
    Upload,
    Uploading,
    Uploaded
  },
  data() {
    return {
      upload: true,
      uploading: false,
      uploaded: false
    }
  },
  methods: {
    changeStatus(payload) {
      if (payload === STATUS.UPLOAD) {
        this.upload = true
        this.uploading = false
        this.uploaded = false
      }

      if (payload === STATUS.UPLOADING) {
        this.upload = false
        this.uploading = true
        this.uploaded = false
      }

      if (payload === STATUS.UPLOADED) {
        this.upload = false
        this.uploading = false
        this.uploaded = true
      }
    }
  }
}
</script>

<style>
.title {
  /* Upload your image */
  position: absolute;
  width: 165px;
  height: 26.99px;
  left: 645.97px;
  top: 340.37px;

  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  line-height: 27px;

  /* identical to box height */
  letter-spacing: -0.035em;

  /* Gray 2 */
  color: #4F4F4F;
}

.description {
  /* File should be Jpeg, Png,... */
  position: absolute;
  width: 121px;
  height: 14.99px;
  left: 660.64px;
  top: 383.35px;

  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 10px;
  line-height: 15px;

  /* identical to box height */
  text-align: center;
  letter-spacing: -0.035em;

  /* Gray 3 */
  color: #828282;
}
</style>

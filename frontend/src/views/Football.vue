<template>
  <div class="profile-page">
    <section class="section-profile-cover section-shaped my-0">
      <div class="shape shape-style-1 shape-primary shape-skew alpha-4">
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
      </div>
    </section>
    <section class="section section-skew">
      <div class="container">
        <modal :show.sync="modals"
               @close="closeModal"
               body-classes="p-0"
               modal-classes="modal-dialog-centered modal-md">
          <card type="secondary" shadow
                header-classes="bg-white pb-12"
                body-classes="px-lg-12 py-lg-12"
                class="border-0">
            <template>
              <form role="form">
                <base-input placeholder="Artist Name" maxlength="200" v-model="formArtistName"></base-input>
                <base-input placeholder="Album Name" maxlength="200" v-model="formAlbumName"></base-input>
                <base-input placeholder="Image Url" maxlength="200" v-model="formImageUrl"></base-input>
                <base-input placeholder="Release Date" v-model="formReleaseDate"></base-input>
                <base-input placeholder="Price" v-model="formPrice"></base-input>
                <base-input placeholder="Sample Url" maxlength="200" v-model="formSampleUrl"></base-input>
                <div class="text-center">
                  <base-button type="success" v-show="isUpdate === false" class="my-4" @click="addMusic()">Create</base-button>
                  <base-button type="success" v-show="isUpdate" class="my-4" @click="updateMusic()">Update</base-button>
                  <base-button type="danger" class="my-4" @click="closeModal()">Close</base-button>
                  <p class="text-success" v-show="errForm === 1 && isUpdate === false">Success add data</p>
                  <p class="text-success" v-show="errForm === 1 && isUpdate === true">Success update data</p>
                  <p class="text-danger" v-show="errForm === 2">Failed to save data</p>
                </div>
              </form>
            </template>
          </card>
        </modal>

        <modal :show.sync="modalsDelete"
               @close="closeModalDelete"
               body-classes="p-0"
               modal-classes="modal-dialog-centered modal-md">
          <card type="secondary" shadow
                header-classes="bg-white pb-12"
                body-classes="px-lg-12 py-lg-12"
                class="border-0">
            <template>
              <div class="text-center" v-show="errFormDelete === 0">
                <p>Are you sure to delete {{this.selectedDeleteName}} ? </p>
                <base-button type="danger" class="my-4" @click="deleteMusic()">Yes</base-button>
                <base-button type="primary" class="my-4" @click="closeModalDelete()">No</base-button>
              </div>
              <div class="text-center">
                <p class="text-success" v-show="errFormDelete === 1">Success to delete data</p>
                <p class="text-danger" v-show="errFormDelete === 2">Failed to delete data</p>
              </div>
              <div class="text-center" v-show="errFormDelete !== 0">
                <base-button type="primary" class="my-4" @click="closeModalDelete()">Close</base-button>
              </div>
            </template>
          </card>
        </modal>

        <card shadow class="card-profile mt--300" no-body>
          <div class="px-4" v-show="loadError">
            <h1>{{this.$root.T("erroroccured")}}</h1>
            <p>{{loadErrorMessage}}</p>
          </div>
          <div class="px-4" v-show="loadError === false">
            <div class="text-center mt-5">
              <h1>Musics</h1>
              <base-button size="sm" type="primary"
                           @click="showCreate()">
                <span class="d-block">Add New</span>
              </base-button>
            </div>
            <b-table striped hover :items="musics" :fields="fields">
              <template #cell(Num)="data">
                {{ data.index + 1 }}
              </template>
              <template #cell(AlbumName)="data">
                <img :src="data.item.ImageUrl" alt="" /> {{data.item.AlbumName}}
              </template>
              <template #cell(SampleUrl)="data">
                <source :src="data.item.SampleUrl" type="audio/mpeg">
              </template>
              <template #cell(Action)="data">
                <base-button size="sm" type="primary"
                             @click="showUpdate(data.item.ArtistId, data.item.ArtistName, data.item.AlbumName, data.item.ImageUrl, data.item.ReleaseDate,data.item.Price, data.item.SampleUrl)">
                  <span class="d-block">Update</span>
                </base-button>
                <base-button size="sm" type="warning"
                             @click="showDelete(data.item.ArtistId, data.item.ArtistName)">
                  <span class="d-block">Delete</span>
                </base-button>
              </template>
            </b-table>
          </div>
        </card>
      </div>
    </section>
  </div>
</template>
<script>
import Modal from "@/components/Modal.vue";
import {calculatePriceNumber} from "@/utils";

export default {
  components: {
    Modal
  },
  data: () => ({
    fields: [{key: 'Num', label: 'Num'},'AlbumName', 'ArtistName', {key: 'ReleaseDate', label: 'Date Release'}, {key: 'SampleUrl', label: 'Sample Audio'}, 'Price', 'Action'],
    musics: [],
    loadError: false,
    loadErrorMessage: '',
    modals: false,
    modalsDelete: false,
    isUpdate: false,

    errForm: 0,
    errFormMessage: '',
    okFormMessage: '',

    errFormDelete: 0,
    errFormDeleteMessage: '',
    okFormDeleteMessage: '',

    formArtistId: 0,
    formArtistName: '',
    formAlbumName: '',
    formImageUrl: '',
    formReleaseDate: '',
    formPrice: 0,
    formSampleUrl: '',

    selectedDeleteName: '',
    selectedDeleteId: 0,
  }),
  mounted() {
    this.loadData(1);
  },
  methods: {
    myTranslation(key) {
      return this.$root.T(key);
    },
    closeModalDelete() {
      this.modalsDelete = false;
    },
    closeModal() {
      this.modals = false;
      this.errForm = 0;

      this.formArtistName = '';
      this.formAlbumName = '';
      this.formImageUrl = '';
      this.formReleaseDate = '';
      this.formPrice = 0;
      this.formSampleUrl = '';
    },
    showCreate() {
      this.isUpdate = false;
      this.modals = true;
    },
    showDelete(id, album_name) {
      this.modalsDelete = true;
      this.errFormDelete = 0;
      this.selectedDeleteId = id;
      this.selectedDeleteName = album_name;
    },
    showUpdate(id, artist_name, album_name, image_url, release_date, price, sample_url) {
      this.modals = true;
      this.isUpdate = true;
      this.setMusicData(id, artist_name, album_name, image_url, release_date, price, sample_url);
    },
    setMusicData(id, artist_name, album_name, image_url, release_date, price, sample_url) {
      this.formArtistId = id;
      this.formArtistName = artist_name;
      this.formAlbumName = album_name;
      this.formImageUrl = image_url;
      this.formReleaseDate = release_date;
      this.formPrice = price;
      this.formSampleUrl = sample_url;
    },
    updateMusic() {
      this.axios.patch('music/'+this.formArtistId, {
        'artist_name': this.formArtistName,
        'album_name': this.formAlbumName,
        'image_url': this.formImageUrl,
        'release_date': this.formReleaseDate,
        'price': parseFloat(this.formPrice),
        'sample_url': this.formSampleUrl,
      }).then((response)=>{
        if (response.data.result) {
          this.errForm = 1;
        } else {
          this.errForm = 2;
        }
      });
    },
    addMusic() {
      this.axios.post('music', {
        'artist_name': this.formArtistName,
        'album_name': this.formAlbumName,
        'image_url': this.formImageUrl,
        'release_date': this.formReleaseDate,
        'price': parseFloat(this.formPrice),
        'sample_url': this.formSampleUrl,
      }).then((response)=>{
        if (response.data.result) {
          this.errForm = 1;
        } else {
          this.errForm = 2;
        }
      });
    },
    deleteMusic() {
      this.axios.delete('music/'+this.selectedDeleteId).then((response)=>{
        if (response.data.result) {
          this.errFormDelete = 1;
        } else {
          this.errFormDelete = 2;
        }
      });
    },
    loadData(p) {
      this.axios.get("/music").then((response) => {
        this.loadError = false;
        if (response.data.result === true) {
          this.musics = response.data.data;
        } else {
          this.loadError = true;
          this.loadErrorMessage = result.message;
        }
      }).catch((err) => {
        this.loadError = true;
        this.loadErrorMessage = this.$root.T("erroroccured");
      });
    },
  },
};
</script>
<style>
</style>

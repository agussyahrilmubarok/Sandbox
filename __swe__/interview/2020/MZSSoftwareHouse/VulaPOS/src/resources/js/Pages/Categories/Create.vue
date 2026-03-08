<script setup>
import { Head, useForm } from '@inertiajs/vue3';
import DashboardLayout from '@/Layouts/DashboardLayout.vue';
import Swal from 'sweetalert2';

defineProps({
    errors: {
        type: Object,
    },
});

const form = useForm({
    name: '',
    description: '',
    image: ''
});

const onSubmit = () => {
    form.post('/dashboard/categories', form, {
        onSuccess: () => {
            Swal.fire({
                title: 'Success!',
                text: 'Category saved successfully.',
                icon: 'success',
                showConfirmButton: false,
                timer: 2000
            });
        }
    });
}

</script>

<template>
    <DashboardLayout>

        <Head title="New Category" />
        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-purple">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-folder"></i> ADD NEW CATEGORY</span>
                                </div>
                                <div class="card-body">

                                    <form @submit.prevent="onSubmit">
                                        <div class="mb-3">
                                            <input class="form-control" @input="form.image = $event.target.files[0]"
                                                :class="{ 'is-invalid': errors.image }" type="file">
                                        </div>
                                        <div v-if="errors.image" class="alert alert-danger">
                                            {{ errors.image }}
                                        </div>
                                        <div class="mb-3">
                                            <label class="fw-bold">Category Name</label>
                                            <input class="form-control" v-model="form.name"
                                                :class="{ 'is-invalid': errors.name }" type="text"
                                                placeholder="Category Name">
                                        </div>
                                        <div v-if="errors.name" class="alert alert-danger">
                                            {{ errors.name }}
                                        </div>
                                        <div class="mb-3">
                                            <label class="fw-bold">Description</label>
                                            <textarea class="form-control" v-model="form.description"
                                                :class="{ 'is-invalid': errors.description }" type="text" rows="4"
                                                placeholder="Description"></textarea>
                                        </div>
                                        <div v-if="errors.description" class="alert alert-danger">
                                            {{ errors.description }}
                                        </div>
                                        <div class="row">
                                            <div class="col-12">
                                                <button class="btn btn-primary shadow-sm rounded-sm"
                                                    type="submit">SAVE</button>
                                                <button class="btn btn-warning shadow-sm rounded-sm ms-3"
                                                    type="reset">RESET</button>
                                            </div>
                                        </div>
                                    </form>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>
    </DashboardLayout>
</template>

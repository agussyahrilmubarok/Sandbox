<script setup>
import { defineProps } from 'vue';
import { Head, useForm } from '@inertiajs/vue3';
import DashboardLayout from '@/Layouts/DashboardLayout.vue';
import Swal from 'sweetalert2';

const props = defineProps({
    errors: {
        type: Object,
    },
    category: {
        type: Object,
    },
});

const form = useForm({
    name: props.category.name,
    description: props.category.description,
    image: ''
});

const onSubmit = () => {
    form.put(`/dashboard/categories/${props.category.id}`, form, {
        onSuccess: () => {
            Swal.fire({
                title: 'Success!',
                text: 'Category updated successfully.',
                icon: 'success',
                showConfirmButton: false,
                timer: 2000
            });
        },
    });
}

</script>

<template>
    <DashboardLayout>

        <Head title="Edit Category" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-purple">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-folder"></i> EDIT CATEGORY</span>
                                </div>
                                <div class="card-body">

                                    <form @submit.prevent="onSubmit">
                                        <div class="mb-3">
                                            <input class="form-control" @input="form.image = $event.target.files[0]"
                                                :class="{ 'is-invalid': errors.imafe }" type="file">
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
                                                    type="submit">UPDATE</button>
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

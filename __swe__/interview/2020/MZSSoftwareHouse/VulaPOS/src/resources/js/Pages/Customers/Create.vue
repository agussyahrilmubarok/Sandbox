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
    phone: '',
    address: ''
});

const onSubmit = () => {
    form.post('/dashboard/customers', form, {
        onSuccess: () => {
            Swal.fire({
                title: 'Success!',
                text: 'Customer saved successfully.',
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

        <Head title="New Customer" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-purple">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-user-circle"></i> ADD NEW
                                        CUSTOMER</span>
                                </div>
                                <div class="card-body">

                                    <form @submit.prevent="onSubmit">
                                        <div class="row">
                                            <div class="col-md-6">
                                                <div class="mb-3">
                                                    <label class="fw-bold">Full Name</label>
                                                    <input class="form-control" v-model="form.name"
                                                        :class="{ 'is-invalid': errors.name }" type="text"
                                                        placeholder="Full Name">
                                                </div>
                                                <div v-if="errors.name" class="alert alert-danger">
                                                    {{ errors.name }}
                                                </div>
                                            </div>
                                            <div class="col-md-6">
                                                <div class="mb-3">
                                                    <label class="fw-bold">Phone</label>
                                                    <input class="form-control" v-model="form.phone"
                                                        :class="{ 'is-invalid': errors.phone }" type="number"
                                                        placeholder="Phone">
                                                </div>
                                                <div v-if="errors.phone" class="alert alert-danger">
                                                    {{ errors.phone }}
                                                </div>
                                            </div>
                                        </div>
                                        <div class="mb-3">
                                            <label class="fw-bold">Address</label>
                                            <textarea class="form-control" v-model="form.address"
                                                :class="{ 'is-invalid': errors.address }" type="text" rows="4"
                                                placeholder="Address"></textarea>
                                        </div>
                                        <div v-if="errors.address" class="alert alert-danger">
                                            {{ errors.address }}
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

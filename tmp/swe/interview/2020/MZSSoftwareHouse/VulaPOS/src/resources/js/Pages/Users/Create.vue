<script setup>
import { Head, useForm } from '@inertiajs/vue3';
import { ref } from 'vue';
import DashboardLayout from '@/Layouts/DashboardLayout.vue';
import Swal from 'sweetalert2';

defineProps({
    errors: {
        type: Object,
    },
    roles: {
        type: Array,
    }
});

const form = useForm({
    name: '',
    email: '',
    password: '',
    password_confirmation: '',
    roles: []
});

const passwordFieldType = ref('password');

const togglePasswordVisibility = () => {
  passwordFieldType.value = passwordFieldType.value === 'password' ? 'text' : 'password';
};

const onSubmit = () => {
    form.post('/dashboard/users', form, {
        onSuccess: () => {
            Swal.fire({
                title: 'Success!',
                text: 'User saved successfully.',
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

        <Head title="New User" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-laravel">

                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-user"></i> ADD NEW USER</span>
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
                                                    <label class="fw-bold">Email Address</label>
                                                    <input class="form-control" v-model="form.email"
                                                        :class="{ 'is-invalid': errors.email }" type="email"
                                                        placeholder="Email Address">
                                                </div>
                                                <div v-if="errors.email" class="alert alert-danger">
                                                    {{ errors.email }}
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-md-6">
                                                <div class="mb-3">
                                                    <label class="fw-bold">Password</label>
                                                    <input class="form-control" v-model="form.password"
                                                        :class="{ 'is-invalid': errors.password }" type="password"
                                                        placeholder="Password">
                                                </div>
                                                <div v-if="errors.password" class="alert alert-danger">
                                                    {{ errors.password }}
                                                </div>
                                            </div>
                                            <div class="col-md-6">
                                                <div class="mb-3">
                                                    <label class="fw-bold">Password Confirmation</label>
                                                    <div class="input-group">
                                                        <input class="form-control" v-model="form.password_confirmation"
                                                            :type="passwordFieldType"
                                                            :class="{ 'is-invalid': errors.password }"
                                                            placeholder="Password" />
                                                        <button class="btn btn-sm btn-outline-secondary"
                                                            @click="togglePasswordVisibility" type="button">
                                                            <i
                                                                :class="passwordFieldType === 'password' ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                                                        </button>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-md-12">
                                                <div class="mb-3">
                                                    <label class="fw-bold">Roles</label>
                                                    <br>
                                                    <div class="form-check form-check-inline"
                                                        v-for="(role, index) in roles" :key="index">
                                                        <input class="form-check-input" type="checkbox"
                                                            v-model="form.roles" :value="role.name"
                                                            :id="`check-${role.id}`">
                                                        <label class="form-check-label" :for="`check-${role.id}`">{{
                                                            role.name }}</label>
                                                    </div>
                                                </div>
                                                <div v-if="errors.roles" class="alert alert-danger">
                                                    {{ errors.roles }}
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-12">
                                                <button class="btn btn-success shadow-sm rounded-sm"
                                                    type="submit">SAVE</button>
                                                <button class="btn btn-secondary shadow-sm rounded-sm ms-3"
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

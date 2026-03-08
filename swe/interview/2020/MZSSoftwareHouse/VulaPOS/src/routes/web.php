<?php

use App\Http\Controllers\PermissionController;
use App\Http\Controllers\ProfileController;
use App\Http\Controllers\Auth\AuthenticatedSessionController;
use App\Http\Controllers\Auth\ConfirmablePasswordController;
use App\Http\Controllers\Auth\EmailVerificationNotificationController;
use App\Http\Controllers\Auth\EmailVerificationPromptController;
use App\Http\Controllers\Auth\NewPasswordController;
use App\Http\Controllers\Auth\PasswordController;
use App\Http\Controllers\Auth\PasswordResetLinkController;
use App\Http\Controllers\Auth\RegisteredUserController;
use App\Http\Controllers\Auth\VerifyEmailController;
use App\Http\Controllers\CategoryController;
use App\Http\Controllers\CustomerController;
use App\Http\Controllers\DashboardController;
use App\Http\Controllers\ProductController;
use App\Http\Controllers\ProfitController;
use App\Http\Controllers\RoleController;
use App\Http\Controllers\SaleController;
use App\Http\Controllers\TransactionController;
use App\Http\Controllers\UserController;
use Illuminate\Foundation\Application;
use Illuminate\Support\Facades\Route;
use Inertia\Inertia;

Route::middleware('guest')->group(function () {
    Route::get('login', [AuthenticatedSessionController::class, 'create'])
        ->name('login');
    Route::post('login', [AuthenticatedSessionController::class, 'store']);
});

Route::prefix('dashboard')->group(function () {
    Route::middleware(['auth', 'verified'])->group(function () {
        // Dashboard
        Route::get('/', [DashboardController::class, 'index'])->name('dashboard');
        // Users
        Route::resource('/users', UserController::class, ['as' => 'dashboard'])
            ->middleware('permission:users.index|users.create|users.edit|users.delete');
        // Roles
        Route::resource('/roles', RoleController::class, ['as' => 'dashboard'])
            ->middleware('permission:roles.index|roles.create|roles.edit|roles.delete');
        // Permissions
        Route::get('/permissions', PermissionController::class)->name('dashboard.permissions.index')
            ->middleware('permission:permissions.index');
        // Categories
        Route::resource('/categories', CategoryController::class, ['as' => 'dashboard'])
            ->middleware('permission:categories.index|categories.create|categories.edit|categories.delete');
        // Products
        Route::resource('/products', ProductController::class, ['as' => 'dashboard'])
            ->middleware('permission:products.index|products.create|products.edit|products.delete');
        // Customers
        Route::resource('/customers', CustomerController::class, ['as' => 'dashboard'])
            ->middleware('permission:customers.index|customers.create|customers.edit|customers.delete');
        // Transactions
        Route::get('/transactions', [TransactionController::class, 'index'])
            ->name('dashboard.transactions.index');
        Route::post('/transactions/searchProduct', [TransactionController::class, 'searchProduct'])
            ->name('dashboard.transactions.searchProduct');
        Route::post('/transactions/addToCart', [TransactionController::class, 'addToCart'])
            ->name('dashboard.transactions.addToCart');
        Route::post('/transactions/destroyCart', [TransactionController::class, 'destroyCart'])
            ->name('dashboard.transactions.destroyCart');
        Route::post('/transactions/store', [TransactionController::class, 'store'])
            ->name('dashboard.transactions.store');
        Route::get('/transactions/print', [TransactionController::class, 'print'])
            ->name('dashboard.transactions.print');
        // Sales
        Route::get('/sales', [SaleController::class, 'index'])->middleware('permission:sales.index')
            ->name('dashboard.sales.index');
        Route::get('/sales/filter', [SaleController::class, 'filter'])
            ->name('dashboard.sales.filter');
        Route::get('/sales/export', [SaleController::class, 'export'])
            ->name('dashboard.sales.export');
        Route::get('/sales/pdf', [SaleController::class, 'pdf'])
            ->name('dashboard.sales.pdf');
        // Profits
        Route::get('/profits', [ProfitController::class, 'index'])
            ->middleware('permission:profits.index')
            ->name('dashboard.profits.index');
        Route::get('/profits/filter', [ProfitController::class, 'filter'])
            ->name('dashboard.profits.filter');
        Route::get('/profits/export', [ProfitController::class, 'export'])
            ->name('dashboard.profits.export');
        Route::get('/profits/pdf', [ProfitController::class, 'pdf'])
            ->name('dashboard.profits.pdf');
    });
});

<?php

namespace Database\Seeders;

use App\Models\Cart;
use App\Models\Category;
use App\Models\Customer;
use App\Models\Product;
use App\Models\Profit;
use App\Models\Transaction;
use App\Models\TransactionDetail;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DummySeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Customer::factory()->count(50)->create();
        Category::factory()->count(50)->create();
        Product::factory()->count(1000)->create();
        Cart::factory()->count(255)->create();
        Transaction::factory()->count(255)->create();
        TransactionDetail::factory()->count(255)->create();
        Profit::factory()->count(255)->create();
    }
}

<?php

namespace App\Http\Controllers;

use App\Models\Customer;
use Illuminate\Http\Request;
use Inertia\Inertia;
use Inertia\Response;

class CustomerController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Request $request): Response
    {
        $q = $request->input('q');

        $customers = Customer::when($q, function ($query, $q) {
            $query->where('name', 'like', '%' . $q . '%');
        })->latest()->paginate(5);

        return Inertia::render('Customers/Index', [
            'customers' => $customers,
        ]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create(): Response
    {
        return Inertia::render('Customers/Create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $request->validate([
            'name' => ['required'],
            'phone' => ['required', 'unique:customers'],
            'address' => ['required'],
        ]);

        Customer::create([
            'name'      => $request->name,
            'phone'   => $request->phone,
            'address'   => $request->address,
        ]);

        return redirect()->route('dashboard.customers.index');
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Customer $customer): Response
    {
        return Inertia::render('Customers/Edit', [
            'customer' => $customer,
        ]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, Customer $customer)
    {
        $request->validate([
            'name' => ['required'],
            'phone' => ['required', 'unique:customers,phone,' . $customer->id],
            'address' => ['required'],
        ]);

        $customer->update([
            'name'      => $request->name,
            'phone'   => $request->phone,
            'address'   => $request->address,
        ]);

        return redirect()->route('dashboard.customers.index');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $customer = Customer::findOrFail($id);
        $customer->delete();

        return redirect()->route('dashboard.customers.index');
    }
}

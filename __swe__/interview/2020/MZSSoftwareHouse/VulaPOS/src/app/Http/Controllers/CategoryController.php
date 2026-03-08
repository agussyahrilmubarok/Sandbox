<?php

namespace App\Http\Controllers;

use App\Models\Category;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;
use Inertia\Inertia;
use Inertia\Response;

class CategoryController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Request $request): Response
    {
        $q = $request->input('q');

        $categories = Category::when($q, function ($query, $q) {
            $query->where('name', 'like', '%' . $q . '%');
        })->latest()->paginate(5);

        return Inertia::render('Categories/Index', [
            'categories' => $categories,
        ]);
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create(): Response
    {
        return Inertia::render('Categories/Create');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $request->validate([
            'name' => ['required', 'unique:categories'],
            'description' => ['required'],
            'image' => ['required', 'image', 'mimes:jpeg,jpg,png', 'max:2000'],
        ]);

        $image = $request->file('image');
        $image->storeAs('public/category', $image->hashName());

        Category::create([
            'name'          => $request->name,
            'description'   => $request->description,
            'image'         => $image->hashName()
        ]);

        return redirect()->route('dashboard.categories.index');
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(Category $category): Response
    {
        return Inertia::render('Categories/Edit', [
            'category' => $category,
        ]);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, Category $category)
    {
        $request->validate([
            'name' => ['required', 'unique:categories,name,'.$category->id],
            'description' => ['required'],
            // 'image' => ['required', 'image', 'mimes:jpeg,jpg,png', 'max:2000'],
        ]);

        //check image update
        if ($request->file('image')) {

            //remove old image
            Storage::disk('local')->delete('public/category/' . basename($category->image));

            //upload new image
            $image = $request->file('image');
            $image->storeAs('public/category', $image->hashName());

            //update category with new image
            $category->update([
                'image' => $image->hashName(),
                'name' => $request->name,
                'description'   => $request->description
            ]);
        }

        $category->update([
            'name'          => $request->name,
            'description'   => $request->description
        ]);

        return redirect()->route('dashboard.categories.index');
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $category = Category::findOrFail($id);
        Storage::disk('local')->delete('public/category/' . basename($category->image));
        $category->delete();

        return redirect()->route('dashboard.categories.index');
    }
}

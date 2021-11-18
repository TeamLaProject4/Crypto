<?php

namespace App\Http\Controllers;

use App\Models\TestTable;
use Illuminate\Http\Request;

class TestTableController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
        return response()->json(TestTable::all());
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
        $data = $request->validate(['text' => 'required']);
        TestTable::create($data);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\TestTable  $TestTable
     * @return \Illuminate\Http\Response
     */
    public function show(TestTable $TestTable)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\TestTable  $TestTable
     * @return \Illuminate\Http\Response
     */
    public function edit(TestTable $TestTable)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\TestTable  $TestTable
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, TestTable $TestTable)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\TestTable  $TestTable
     * @return \Illuminate\Http\Response
     */
    public function destroy(TestTable $TestTable)
    {
        //
    }
}

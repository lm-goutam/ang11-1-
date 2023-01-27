import { Component } from '@angular/core';
import { read, utils, writeFile } from 'xlsx';

@Component({
  selector: 'app-table-view',
  templateUrl: './table-view.component.html',
  styleUrls: ['./table-view.component.css']
})
export class TableViewComponent {
td: any;
    post: any;


   
  title(_title: any) {
    throw new Error('Method not implemented.');
  }
  
    movies: any[] = [];
user: any;

    handleImport($event: any) {
        const files = $event.target.files;
        if (files.length) {
            const file = files[0];
            const reader = new FileReader();
            reader.onload = (event: any) => {
                const wb = read(event.target.result);
                const sheets = wb.SheetNames;

                if (sheets.length) {
                    const rows = utils.sheet_to_json(wb.Sheets[sheets[0]]);
                    this.movies = rows;
                }
            }
            reader.readAsArrayBuffer(file);
        }
    }

    handleExport() {
        const headings = [[
            // 'id',
            // 'Name',
            // 'Address',
            // 'Salary',
            // 'UserName'
        ]];
        const wb = utils.book_new();
        const ws: any = utils.json_to_sheet([]);
        utils.sheet_add_aoa(ws, headings);
        utils.sheet_add_json(ws, this.movies, { origin: 'A2', skipHeader: true });
        utils.book_append_sheet(wb, ws, 'Report');
        writeFile(wb, 'Movie Report.xlsx');
    }
    th:any;
  data:any=[];
  getData(){
    this.post.getData().subscribe(
      (data: any) =>{
        this.data=data;
     
      }
    );
  }
}



// doctor.component.ts
import { Component } from '@angular/core';

@Component({
  selector: 'app-doctor',
  templateUrl: './doctor.component.html',
  styleUrls: ['./doctor.component.css']
})
export class DoctorComponent {
  data = [
    { date: "11/2/2023", time: '11 pm', doctor: "ahmed" },
    { date: "11/4/2023", time: '10 am', doctor: "sakr" },
    { date: "11/7/2023", time: '6 pm', doctor: "wael" }
  ];

  newSlot: { date: string, time: string, doctor: string } = { date: '', time: '', doctor: '' };
  editingIndex: number | null = null;

  addNewSlot() {
    if (this.newSlot.date && this.newSlot.time) {
      this.data.unshift({ date: this.newSlot.date, time: this.newSlot.time, doctor: this.newSlot.doctor });
      this.newSlot.date = '';
      this.newSlot.time = '';
    }
  }

  editSlot(index: number) {
    this.editingIndex = index;
  }

  saveEdit(index: number) {
    // Save the edited values and clear the editing state
    this.data[index].date = this.newSlot.date;
    this.data[index].time = this.newSlot.time;
    this.editingIndex = null;
  }

  cancelEdit() {
    // Cancel the editing operation and clear the editing state
    this.editingIndex = null;
  }

  cancelSlot(index: number) {
    this.data.splice(index, 1);
  }
}

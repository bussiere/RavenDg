from django.db import models
from RatePayer.models import TaxUser



class FileCase(models.Model):
    data = models.FileField(blank=True, null=True)

# Create your models here.
class CategoryCase(models.Model):
    name = models.CharField(blank=True, null=True, max_length=500)



class Case(models.Model):
    category = models.ForeignKey(CategoryCase,blank=True, null=True)
    user = models.ForeignKey(TaxUser,blank=True, null=True)
    note = models.CharField(blank=True, null=True, max_length=500)
    dateOpen = models.DateField(blank=True, null=True, max_length=200)
    dateClose = models.DateField(blank=True, null=True, max_length=200)
    file = models.ManyToManyField(FileCase)
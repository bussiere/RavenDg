from django.db import models
from RatePayer.models import TaxUser


# Create your models here.
class CategoryCase(models.Model):
    name = models.CharField(blank=True, null=True, max_length=500)



class Case(models.Model):
    category = models.ForeignKey(CategoryCase,blank=True, null=True,on_delete=models.PROTECT)
    user = models.ForeignKey(TaxUser,blank=True, null=True,on_delete=models.PROTECT,related_name="UserLitigation")
    note = models.CharField(blank=True, null=True, max_length=500)
    dateOpen = models.DateField(blank=True, null=True, max_length=200)
    dateClose = models.DateField(blank=True, null=True, max_length=200)